

//////////////////////////////////////////////////////////////////////////
//
// Definitions of error types.
//

export interface GenericError<Category extends string> {
	success: false;
	category: Category;
	message: string;
}

export type DispatchError = GenericError<'dispatch'>
export type ValidationError = GenericError<'validation'>
export type InternalError = GenericError<'internal'>
export type ExecutionError = GenericError<'execution'>


//////////////////////////////////////////////////////////////////////////
//
// Factories for error types.
//

export function createDispatchError(message: string): DispatchError {
	return { success: false, category: 'dispatch' as const, message };
}

export function createValidationError(message: string): ValidationError {
	return { success: false, category: 'validation' as const, message };
}

export function createInternalError(message: string): InternalError {
	return { success: false, category: 'internal' as const, message };
}

export function createExecutionError(input: any): ExecutionError {
	if (typeof input === 'object' && input instanceof Error) {
		return { success: false, category: 'execution' as const, message: input.message };
	} else if (typeof input === 'string') {
		return { success: false, category: 'execution' as const, message: input };
	} else {
		return { success: false, category: 'execution' as const, message: 'unknown error' };
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Type guards for error types.
//

export function isDispatchError(input: any): input is DispatchError {
	if (isBaseError(input) === false) {
		return false;
	} else if (input.category !== 'dispatch') {
		return false;
	} else {
		return true;
	}
}

export function isValidationError(input: any): input is ValidationError {
	if (isBaseError(input) === false) {
		return false;
	} else if (input.category !== 'validation') {
		return false;
	} else {
		return true;
	}
}

export function isInternalError(input: any): input is InternalError {
	if (isBaseError(input) === false) {
		return false;
	} else if (input.category !== 'internal') {
		return false;
	} else {
		return true;
	}
}

export function isExecutionError(input: any): input is ExecutionError {
	if (isBaseError(input) === false) {
		return false;
	} else if (input.category !== 'execution') {
		return false;
	} else {
		return true;
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Helper types and type guards.
//

interface BaseError {
	success: false;
	category: string;
	message: string;
}

function isBaseError(input: any): input is BaseError {
	if (typeof input !== 'object') {
		return false;
	} else if (typeof input.success !== 'boolean') {
		return false;
	} else if (typeof input.category !== 'string') {
		return false;
	} else if (typeof input.message !== 'string') {
		return false;
	} else if (input.success !== false) {
		return false;
	} else {
		return true;
	}
}


