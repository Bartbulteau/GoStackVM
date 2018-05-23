call myprint 0
jmp end

myprint:
    const_i32 10
    const_i32 20
    add_i32
    print_i32
    
    const_i32 0 ret

end: halt