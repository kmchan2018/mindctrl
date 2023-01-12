

import * as WebExtension from 'webextension-polyfill';
import * as Pattern from './pattern';


//////////////////////////////////////////////////////////////////////////
//
// Validate the type of the input against the given type guards. The
// typing of the function is more complex than usual, so here is the
// explanation.
//
// The GT generic type stands for guard tuple. The types of its members
// mirrors the types of type guards passed into the function.
//
// The GT[number] extracts the types of members in GT and combines them
// into a union type.
//
// Then, the distributive conditional type part extracts the target
// type of each type guards in GT and combines them into a union type.
//
// The only problem with the typing is that the function can accept
// functions that are not type guards. In such case, the function
// continues to works but cannot narrow the type of input.
//

export function validateType<GT extends Array<(input: any) => boolean>>(input: any, ...guards: GT): input is (GT[number] extends (input: any) => input is infer T ? T : any) {
	return guards.some((guard) => guard(input));
}


//////////////////////////////////////////////////////////////////////////
//
// Type guards for basic javascript types.
//

export function isBoolean(input: any): input is boolean {
	return typeof input === 'boolean';
}

export function isNumber(input: any): input is number {
	return typeof input === 'number';
}

export function isString(input: any): input is string {
	return typeof input === 'string';
}

export function isRecord(input: any): input is Record<string,any> {
	return typeof input === 'object';
}

export function isNull(input: any): input is null {
	return input === null;
}

export function isUndefined(input: any): input is undefined {
	return input === undefined;
}

export function isTrue(input: any): input is true {
	return input === true;
}

export function isFalse(input: any): input is false {
	return input === false;
}

export function isLiteral<T>(literal: T): (input: any) => input is T {
	return function(input: any): input is T {
		return input === literal;
	};
}


//////////////////////////////////////////////////////////////////////////
//
// Type guards for literal types.
//

export function isAll(input: any): input is 'all' {
	return input === 'all';
}

export function isCurrent(input: any): input is 'current' {
	return input === 'all';
}


//////////////////////////////////////////////////////////////////////////
//
// Type guards for specialized string types.
//

export function isUrl(input: any): input is string {
	if (typeof input !== 'string') {
		return false;
	} else if (input === '' || input.trim() === '') {
		return false;
	} else {
		return true;
	}
}

export function isFilename(input: any): input is string {
	if (typeof input !== 'string') {
		return false;
	} else if (input === '' || input.trim() === '') {
		return false;
	} else {
		return true;
	}
}

export function isMatchPattern(input: any): input is string {
	if (typeof input !== 'string') {
		return false;
	} else if (Pattern.validateMatchPattern(input) === false) {
		return false;
	} else {
		return true;
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Type guards for web extension types.
//

export function isDownloadId(input: any): input is number {
	return typeof input === 'number';
}

export function isDownloadState(input: any): input is WebExtension.Downloads.State {
	if (typeof input !== 'string') {
		return false;
	} else if (input !== 'in_progress' && input !== 'interrupted' && input !== 'complete') {
		return false;
	} else {
		return true;
	}
}

export function isTabId(input: any): input is number {
	return typeof input === 'number';
}

export function isTabStatus(input: any): input is WebExtension.Tabs.TabStatus {
	if (typeof input !== 'string') {
		return false;
	} else if (input !== 'loading' && input !== 'complete') {
		return false;
	} else {
		return true;
	}
}

export function isWindowId(input: any): input is number {
	return typeof input === 'number';
}

export function isCssOrigin(input: any): input is WebExtension.ExtensionTypes.CSSOrigin {
	if (input === 'author') {
		return true;
	} else if (input === 'user') {
		return true;
	} else {
		return false;
	}
}

