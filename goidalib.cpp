#include "goidalib.hpp"

// Implementation of IDA library functions

int IDALib::Init() {
    return init_library(0, nullptr);
}

int IDALib::OpenDatabase(const char *file_path, bool run_auto) {
    return open_database(file_path, run_auto);
}

void IDALib::CloseDatabase(bool save) {
    close_database(save);
}

bool IDALib::MakeSignatures(bool only_pat) {
    return make_signatures(only_pat);
}

void IDALib::EnableConsoleMessages(bool enable) {
    enable_console_messages(enable);
}

void IDALib::SetScreenEA(ea_t ea) {
    set_screen_ea(ea);
}

bool IDALib::GetLibraryVersion(int &major, int &minor, int &build) {
    return get_library_version(major, minor, build);
}

IDALib::IDALib() {
    // Constructor implementation
}

IDALib::~IDALib() {
    // Destructor implementation
}