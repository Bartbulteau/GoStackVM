const_i32 100000
gstore 0
call myprint 0
jmp end

myprint:
    const_i32 200
    store 0
    gload 0
    print_i32
    load 0
    print_i32


end: halt