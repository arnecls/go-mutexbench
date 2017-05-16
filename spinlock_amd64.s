#include "textflag.h"

// func Lock32(v *int32) bool
TEXT ·Lock32(SB), NOSPLIT, $0-16
    MOVQ    v(FP), BX
    MOVL    $1, AX          // prepare return value
    LOCK 
    BTSQ    $0, (BX)        // store previous value of bit 0 in CF (carry)
    SBBL    $0, AX          // store carry as result: AX - (CF+0)
    MOVL    AX, ret+8(FP)
    RET

// func Unlock32(v *int32)
TEXT ·Unlock32(SB), NOSPLIT, $0-8
    MOVQ    v(FP), BX
    LOCK 
    BTRQ    $0, (BX)        // Set bit 0 to 0
    RET
