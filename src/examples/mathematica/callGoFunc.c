#include "WolframLibrary.h"
#include "libgoSquare.h"

DLLEXPORT int callGoFunc(WolframLibraryData libData, mint Argc, MArgument *Args, MArgument Res) {
	mint in;
	in = MArgument_getInteger(Args[0]);
	int res = GoSquare((int)in);
	MArgument_setInteger(Res, res);
	return LIBRARY_NO_ERROR;
}