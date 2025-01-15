#include <string>
#include <pro.h>
#include <ida.hpp>
#include <idalib.hpp>
#include <funcs.hpp>
#include <hexrays.hpp>

#ifndef _GOIDALIB_H_
#define _GOIDALIB_H_

class IDALib {
public:
    IDALib();
    ~IDALib();
    int Init();
    int OpenDatabase(const char *file_path, bool run_auto);
    void CloseDatabase(bool save);
    bool MakeSignatures(bool only_pat);
    void EnableConsoleMessages(bool enable);
    void SetScreenEA(ea_t ea);
    // func_t *GetFunc(ea_t ea);
    func_t *GetFuncByID(int id);
    int FuncCount();
    func_t *GetFunc(uint64 ea);
    bool GetLibraryVersion(int &major, int &minor, int &build);
    std::string DecompileFunction(func_t *pfn);
};

#endif // _GOIDALIB_H_