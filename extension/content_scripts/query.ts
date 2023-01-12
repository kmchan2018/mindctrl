

import { GraphQLSchema, GraphQLObjectType, GraphQLString, GraphQLNonNull, GraphQLList, graphql } from 'graphql';


//////////////////////////////////////////////////////////////////////////
//
// GraphQL value object for elements.
//

class ElementValue {
	private target: HTMLElement;

	constructor(element: HTMLElement) {
		this.target = element;
	}

	me(): ElementValue {
		return this;
	}

	html(): string {
		return this.target.innerHTML;
	}

	text(): string | null {
		return this.target.textContent;
	}

	content(): string | null {
		return this.target.innerText;
	}

	style(args: { name: string }) {
		return this.target.style.getPropertyValue(args.name);
	}

	attribute(args: { name: string }): string | null {
		return this.target.getAttribute(args.name);
	}

	urlattribute(args: { name: string; component?: string }): string | null {
		const value = this.target.getAttribute(args.name);
		
		if (value) {
			const url = new URL(value, document.URL);

			switch (args.component) {
				case 'basename':  return url.pathname.split('/').pop() || null;
				case 'host':      return url.host;
				case 'hostname':  return url.hostname;
				case 'href':      return url.href;
				case 'password':  return url.password;
				case 'pathname':  return url.pathname;
				case 'port':      return url.port;
				case 'protocol':  return url.protocol;
				case 'search':    return url.search;
				case 'username':  return url.username;
				case undefined:   return url.href;
				default:          return url.href;
			}
		} else {
			return null;
		}
	}

	element(args: { selector: string }) {
		const selection = this.target.querySelector<HTMLElement>(args.selector);
		if (selection) {
			return new ElementValue(selection);
		} else {
			return undefined;
		}
	}

	elements(args: { selector: string }) {
		return Array.
			from(this.target.querySelectorAll<HTMLElement>(args.selector)).
			map((element) => new ElementValue(element));
	}
}


//////////////////////////////////////////////////////////////////////////
//
// GraphQL value object for the document.
//

class DocumentValue {
	url() {
		return document.URL;
	}

	title() {
		return document.title;
	}

	referrer() {
		return document.referrer;
	}

	root() {
		return new ElementValue(document.documentElement);
	}

	element(args: { selector: string }) {
		const selection = document.querySelector<HTMLElement>(args.selector);
		if (selection) {
			return new ElementValue(selection);
		} else {
			return undefined;
		}
	}

	elements(args: { selector: string }) {
		return Array.
			from(document.querySelectorAll<HTMLElement>(args.selector)).
			map((element) => new ElementValue(element));
	}
}


//////////////////////////////////////////////////////////////////////////
//
// GraphQL object type for elements.
//

const ElementDefinition: GraphQLObjectType = new GraphQLObjectType({
	name: "Element",
	description: "GraphQL type representing a HTMLElement object in the document",
    
	fields: () => ({
		me: {
			description: "Element itself",
			type: new GraphQLNonNull(ElementDefinition),
		},

		html: {
			description: "HTML code of the element",
			type: GraphQLString,
		},

		text: {
			description: "Textual data under the element",
			type: GraphQLString,
		},

		content: {
			description: "User visible text under the element",
			type: GraphQLString,
		},

		style: {
			description: "Value of the given CSS property",
			type: GraphQLString,
			args: {
				name: {
					description: "name of the style",
					type: GraphQLString,
				},
			},
		},

		attribute: {
			description: "Value of the given attribute",
			type: GraphQLString,
			args: {
				name: {
					description: "name of the attribute",
					type: GraphQLString,
				},
			},
		},

		urlattribute: {
			description: "URL component of the value of the given attribute",
			type: GraphQLString,
			args: {
				name: {
					description: "name of the attribute",
					type: GraphQLString,
				},
				component: {
					description: "url component",
					type: GraphQLString,
				},
			},
		},

		element: {
			description: "first descendant that matches the specified selector",
			type: ElementDefinition,
			args: {
				selector: {
					description: "selector used to filter the descendants",
					type: new GraphQLNonNull(GraphQLString),
				}
			}
		},

		elements: {
			description: "all descendants that matches the specified selector",
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(ElementDefinition))),
			args: {
				selector: {
					description: "selector used to filter the descendants",
					type: new GraphQLNonNull(GraphQLString),
				}
			}
		},
	}),
});


//////////////////////////////////////////////////////////////////////////
//
// GraphQL object type for the document.
//

const DocumentDefinition: GraphQLObjectType = new GraphQLObjectType({
	name: "Document",
	description: "GraphQL type representing the document",

	fields: () => ({
		url: {
			description: "URL of the document",
			type: new GraphQLNonNull(GraphQLString),
		},

		title: {
			description: "Title of the document",
			type: new GraphQLNonNull(GraphQLString),
		},

		referrer: {
			description: "Referrer of the document",
			type: new GraphQLNonNull(GraphQLString),
		},

		root: {
			description: "Root element of the document",
			type: new GraphQLNonNull(ElementDefinition),
		},

		element: {
			description: "first element in the document that matches the specified selector",
			type: ElementDefinition,
			args: {
				selector: {
					description: "selector used to filter the elements",
					type: new GraphQLNonNull(GraphQLString),
				}
			}
		},

		elements: {
			description: "all elements in the document that matches the specified selector",
			type: new GraphQLNonNull(new GraphQLList(new GraphQLNonNull(ElementDefinition))),
			args: {
				selector: {
					description: "selector used to filter the elements",
					type: new GraphQLNonNull(GraphQLString),
				}
			}
		},
	}),
});


//////////////////////////////////////////////////////////////////////////
//
// GraphQL root schema.
//

const Schema = new GraphQLSchema({
	query: DocumentDefinition,
});


//////////////////////////////////////////////////////////////////////////
//
// GraphQL root value object.
//

const Value = new DocumentValue();


//////////////////////////////////////////////////////////////////////////
//
// Public API.
//

interface QueryResult {
	success: true;
	result: any;
}

interface QueryError {
	success: false;
	origin: 'content_script';
	message: string;
}

export async function Invoke(query: string, operation?: string, variables?: Record<string,any>): Promise<QueryResult|QueryError> {
	try {
		const result = await graphql({
			schema: Schema, 
			rootValue: Value,
			source: query,
			variableValues: variables,
			operationName: operation,
		});

		if (result.errors) {
			return { success: false, origin: 'content_script', message: result.errors[0].message };
		} else {
			return { success: true, result: result.data };
		}
	} catch (error) {
		if (error instanceof Error) {
			return { success: false, origin: 'content_script', message: error.message };
		} else if (typeof error === 'string') {
			return { success: false, origin: 'content_script', message: error };
		} else {
			console.log('[BUG] value cannot be converted to string for reporting: %s', error);
			return { success: false, origin: 'content_script', message: 'unknown error' };
		}
	}
}


