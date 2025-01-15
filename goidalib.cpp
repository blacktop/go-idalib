#include <iostream>
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

int IDALib::FuncCount() {
    return get_func_qty();
}

func_t *IDALib::GetFuncByID(int id) {
    return getn_func(id);
}

func_t *IDALib::GetFunc(uint64 ea) {
    return get_func(ea_t(ea));
}

std::string IDALib::DecompileFunction(func_t *pfn) {
  hexrays_failure_t hf;
  cfuncptr_t cfunc = decompile(pfn, &hf, DECOMP_WARNINGS);
  if ( cfunc == nullptr )
  {
    warning("#error \"%a: %s", hf.errea, hf.desc().c_str());
    return "#error " + std::to_string(hf.errea) + ": " + hf.desc().c_str();
  }
  msg("%a: successfully decompiled\n", pfn->start_ea);

  std::string decompiled_code;
  const strvec_t &sv = cfunc->get_pseudocode();
  for ( int i=0; i < sv.size(); i++ )
  {
    qstring buf;
    tag_remove(&buf, sv[i].line);
    decompiled_code += buf.c_str();
    decompiled_code += "\n";
  }

  return decompiled_code;
}

IDALib::IDALib() {
    // Constructor implementation
}

IDALib::~IDALib() {
    // Destructor implementation
}