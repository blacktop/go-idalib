#include <pro.h>
#include <ida.hpp>
#include <idalib.hpp>

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
    bool GetLibraryVersion(int &major, int &minor, int &build);
};

#endif // _GOIDALIB_H_