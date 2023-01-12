

//////////////////////////////////////////////////////////////////////////
//
// Regular expression used for parsing match patterns. Match patterns are
// text patterns web extension used to match an URL. A description of the
// syntax can be found at:
//
// https://developer.chrome.com/docs/extensions/mv3/match_patterns/
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Match_patterns
//
// Note that the special <all_urls> sentinel value is NOT supported by
// this module.
//

const MATCH_PATTERN_EXTRACTOR = /^(http|https|ws|wss|ftp|data|file|\*):\/\/((?:(?:\*|[A-Za-z0-9\x2d]+)(?:\x2e[A-Za-z0-9\x2d]+)*)?)(\/[^?]*)(\?.*)?$/i;


//////////////////////////////////////////////////////////////////////////
//
// Escape lookup table for ASCII characters. It is used by string escape
// functions to determine if an ASCII character should be escaped inside
// regular expression.
//

const REGEXP_WHITELIST = (function(): Array<boolean> {
	const whitelist = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_#&%@";
	const output = [];

	for (let i = 0; i < 256; i++) {
		output.push(whitelist.includes(String.fromCharCode(i)));
	}

	return output;
})();


//////////////////////////////////////////////////////////////////////////
//
// Validate a match pattern. The validation is done by running two
// checks:
//
// First of all, the pattern is matched against the extractor regular
// expression. A match ensures that the pattern follows the required
// syntax.
//
// Then, the scheme and host are checked to make sure that no scheme
// other than file has an empty host.
//

export function validateMatchPattern(pattern: string): boolean {
	const parts = MATCH_PATTERN_EXTRACTOR.exec(pattern);

	if (parts) {
		const scheme = parts[1] as string;
		const host = parts[2] as string;

		if (host === '' && scheme !== 'file') {
			return false;
		} else {
			return true;
		}
	} else {
		return false;
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Convert a match pattern to the corresponding regular expression.
//

export function convertMatchPattern(pattern: string): string {
	const parts = MATCH_PATTERN_EXTRACTOR.exec(pattern);

	if (parts) {
		const scheme = parts[1] as string;
		const host = parts[2] as string;
		const path = parts[3] as string;
		const query = parts[4] as string;
		const output = ['^'];

		if (scheme !== '*') {
			output.push(scheme);
			output.push('\\x3a\\x2f\\x2f' /* "://" */);
		} else {
			output.push('(http|https)');
			output.push('\\x3a\\x2f\\x2f' /* "://" */);
		}

		if (host === '*') {
			output.push('[A-Za-z0-9_\\x2d\\x2e]+');
		} else if (host.startsWith('*.') === true) {
			output.push('[A-Za-z0-9_\\x2d\\x2e]+');
			output.push('\\x2e' /* "." */);
			convertLiteralStringInto(host.substring(2), output);
		} else if (host !== '') {
			convertLiteralStringInto(host, output);
		}

		if (query) {
			convertWildcardStringInto(path, '[^\\?]*', output);
			convertWildcardStringInto(query, '.*', output);
			output.push('$');
		} else {
			convertWildcardStringInto(path, '[^\\?]*', output);
			output.push('$');
		}

		return output.join('');
	} else {
		throw new Error('invalid match pattern');
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Translate a literal string to the corresponding regular expression
// fragnment.
//

function convertLiteralStringInto(literal: string, output: Array<string>) {
	for (let i = 0; i < literal.length; i++) {
		escapeCodeUnitInto(literal.charCodeAt(i)!, output);
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Translate a wildcard string to the corresponding regular expression
// fragnment.
//

function convertWildcardStringInto(wildcard: string, subst: string, output: Array<string>) {
	for (let i = 0; i < wildcard.length; i++) {
		const charcode = wildcard.charCodeAt(i)!;

		if (charcode !== 0x2a /* "*" */) {
			escapeCodeUnitInto(charcode, output);
		} else {
			output.push(subst);
		}
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Generate a regular expression fragment that matches and only matches
// the given UTF-16 code unit.
//

function escapeCodeUnitInto(codeunit: number, output: Array<string>) {
	if (codeunit < 0) {
		throw new RangeError('code unit out of range');
	} else if (codeunit > 0xffff) {
		throw new RangeError('code unit out of range');
	}

	if (REGEXP_WHITELIST[codeunit] === true) {
		output.push(String.fromCodePoint(codeunit));
	} else if (codeunit <= 0xf) {    // \x0#
		output.push('\\x0');
		output.push(codeunit.toString(16));
	} else if (codeunit <= 0xff) {   // \x##
		output.push('\\x');
		output.push(codeunit.toString(16));
	} else if (codeunit <= 0xfff) {  // \u0###
		output.push('\\u0');
		output.push(codeunit.toString(16));
	} else {                         // \u####
		output.push('\\u');
		output.push(codeunit.toString(16));
	}
}


