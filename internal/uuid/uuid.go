//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package uuid ;import (_fd "crypto/rand";_a "encoding/hex";_g "io";);var _ff =_fd .Reader ;func (_af UUID )String ()string {var _fb [36]byte ;_bb (_fb [:],_af );return string (_fb [:])};type UUID [16]byte ;var Nil =_aff ;var _aff UUID ;func MustUUID ()UUID {uuid ,_b :=NewUUID ();
if _b !=nil {panic (_b );};return uuid ;};func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_e :=_g .ReadFull (_ff ,uuid [:]);if _e !=nil {return _aff ,_e ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};func _bb (_eb []byte ,_afb UUID ){_a .Encode (_eb ,_afb [:4]);
_eb [8]='-';_a .Encode (_eb [9:13],_afb [4:6]);_eb [13]='-';_a .Encode (_eb [14:18],_afb [6:8]);_eb [18]='-';_a .Encode (_eb [19:23],_afb [8:10]);_eb [23]='-';_a .Encode (_eb [24:],_afb [10:]);};