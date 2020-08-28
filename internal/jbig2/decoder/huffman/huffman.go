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

package huffman ;import (_e "errors";_gc "fmt";_b "github.com/unidoc/unipdf/v3/internal/bitwise";_gf "math";_f "strings";);func (_gfd *EncodedTable )RootNode ()*InternalNode {return _gfd ._d };func _ee (_fcd int32 )*InternalNode {return &InternalNode {_gcb :_fcd }};func (_edgd *InternalNode )pad (_eb *_f .Builder ){for _bdc :=int32 (0);_bdc < _edgd ._gcb ;_bdc ++{_eb .WriteString ("\u0020\u0020\u0020");};};func (_ge *FixedSizeTable )Decode (r _b .StreamReader )(int64 ,error ){return _ge ._dg .Decode (r )};func _eee (_ccf ,_dafb int32 )int32 {if _ccf > _dafb {return _ccf ;};return _dafb ;};func (_gab *StandardTable )Decode (r _b .StreamReader )(int64 ,error ){return _gab ._dcb .Decode (r )};type EncodedTable struct{BasicTabler ;_d *InternalNode ;};func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_gbc :prefixLength ,_bdb :rangeLength ,_eec :rangeLow ,_cca :isLowerRange ,_ccc :-1};};func _ec (_cce *Code )*OutOfBandNode {return &OutOfBandNode {}};type InternalNode struct{_gcb int32 ;_cdd Node ;_gce Node ;};func (_cga *ValueNode )String ()string {return _gc .Sprintf ("\u0025\u0064\u002f%\u0064",_cga ._ad ,_cga ._eda );};func _ccb (_aa []*Code ){var _abe int32 ;for _ ,_edc :=range _aa {_abe =_eee (_abe ,_edc ._gbc );};_agb :=make ([]int32 ,_abe +1);for _ ,_ebf :=range _aa {_agb [_ebf ._gbc ]++;};var _ggb int32 ;_eg :=make ([]int32 ,len (_agb )+1);_agb [0]=0;for _fcef :=int32 (1);_fcef <=int32 (len (_agb ));_fcef ++{_eg [_fcef ]=(_eg [_fcef -1]+(_agb [_fcef -1]))<<1;_ggb =_eg [_fcef ];for _ ,_bcc :=range _aa {if _bcc ._gbc ==_fcef {_bcc ._ccc =_ggb ;_ggb ++;};};};};type ValueNode struct{_ad int32 ;_eda int32 ;_ca bool ;};func (_fae *StandardTable )RootNode ()*InternalNode {return _fae ._dcb };var _ Tabler =&EncodedTable {};func (_gd *EncodedTable )InitTree (codeTable []*Code )error {_ccb (codeTable );for _ ,_df :=range codeTable {if _cf :=_gd ._d .append (_df );_cf !=nil {return _cf ;};};return nil ;};func (_fff *EncodedTable )parseTable ()error {var (_bf []*Code ;_fe ,_da ,_bd int32 ;_fd uint64 ;_cc error ;);_cg :=_fff .StreamReader ();_gge :=_fff .HtLow ();for _gge < _fff .HtHigh (){_fd ,_cc =_cg .ReadBits (byte (_fff .HtPS ()));if _cc !=nil {return _cc ;};_fe =int32 (_fd );_fd ,_cc =_cg .ReadBits (byte (_fff .HtRS ()));if _cc !=nil {return _cc ;};_da =int32 (_fd );_bf =append (_bf ,NewCode (_fe ,_da ,_bd ,false ));_gge +=1<<uint (_da );};_fd ,_cc =_cg .ReadBits (byte (_fff .HtPS ()));if _cc !=nil {return _cc ;};_fe =int32 (_fd );_da =32;_bd =_fff .HtLow ()-1;_bf =append (_bf ,NewCode (_fe ,_da ,_bd ,true ));_fd ,_cc =_cg .ReadBits (byte (_fff .HtPS ()));if _cc !=nil {return _cc ;};_fe =int32 (_fd );_da =32;_bd =_fff .HtHigh ();_bf =append (_bf ,NewCode (_fe ,_da ,_bd ,false ));if _fff .HtOOB ()==1{_fd ,_cc =_cg .ReadBits (byte (_fff .HtPS ()));if _cc !=nil {return _cc ;};_fe =int32 (_fd );_bf =append (_bf ,NewCode (_fe ,-1,-1,false ));};if _cc =_fff .InitTree (_bf );_cc !=nil {return _cc ;};return nil ;};func _bcb (_cd *Code )*ValueNode {return &ValueNode {_ad :_cd ._bdb ,_eda :_cd ._eec ,_ca :_cd ._cca }};func _fce (_fee [][]int32 )(*StandardTable ,error ){var _dcg []*Code ;for _cbd :=0;_cbd < len (_fee );_cbd ++{_db :=_fee [_cbd ][0];_faf :=_fee [_cbd ][1];_dgd :=_fee [_cbd ][2];var _cdg bool ;if len (_fee [_cbd ])> 3{_cdg =true ;};_dcg =append (_dcg ,NewCode (_db ,_faf ,_dgd ,_cdg ));};_fb :=&StandardTable {_dcb :_ee (0)};if _af :=_fb .InitTree (_dcg );_af !=nil {return nil ,_af ;};return _fb ,nil ;};func _eab (_dfa ,_cdf int32 )string {var _ac int32 ;_bg :=make ([]rune ,_cdf );for _cgb :=int32 (1);_cgb <=_cdf ;_cgb ++{_ac =_dfa >>uint (_cdf -_cgb )&1;if _ac !=0{_bg [_cgb -1]='1';}else {_bg [_cgb -1]='0';};};return string (_bg );};func (_dff *OutOfBandNode )Decode (r _b .StreamReader )(int64 ,error ){return int64 (_gf .MaxInt64 ),nil };type Node interface{Decode (_ed _b .StreamReader )(int64 ,error );String ()string ;};var _ Node =&OutOfBandNode {};var _gcc =make ([]Tabler ,len (_ab ));func (_fc *OutOfBandNode )String ()string {return _gc .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_gf .MaxInt64 ));};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_a :=&EncodedTable {_d :&InternalNode {},BasicTabler :table };if _gg :=_a .parseTable ();_gg !=nil {return nil ,_gg ;};return _a ,nil ;};func (_fad *Code )String ()string {var _ebgb string ;if _fad ._ccc !=-1{_ebgb =_eab (_fad ._ccc ,_fad ._gbc );}else {_ebgb ="\u003f";};return _gc .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_ebgb ,_fad ._gbc ,_fad ._bdb ,_fad ._eec );};func (_daf *InternalNode )append (_ga *Code )(_fca error ){if _ga ._gbc ==0{return nil ;};_ae :=_ga ._gbc -1-_daf ._gcb ;if _ae < 0{return _e .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");};_fa :=(_ga ._ccc >>uint (_ae ))&0x1;if _ae ==0{if _ga ._bdb ==-1{if _fa ==1{if _daf ._gce !=nil {return _gc .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_ga );};_daf ._gce =_ec (_ga );}else {if _daf ._cdd !=nil {return _gc .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_ga );};_daf ._cdd =_ec (_ga );};}else {if _fa ==1{if _daf ._gce !=nil {return _gc .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_ga );};_daf ._gce =_bcb (_ga );}else {if _daf ._cdd !=nil {return _gc .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_ga );};_daf ._cdd =_bcb (_ga );};};}else {if _fa ==1{if _daf ._gce ==nil {_daf ._gce =_ee (_daf ._gcb +1);};if _fca =_daf ._gce .(*InternalNode ).append (_ga );_fca !=nil {return _fca ;};}else {if _daf ._cdd ==nil {_daf ._cdd =_ee (_daf ._gcb +1);};if _fca =_daf ._cdd .(*InternalNode ).append (_ga );_fca !=nil {return _fca ;};};};return nil ;};func (_agc *FixedSizeTable )InitTree (codeTable []*Code )error {_ccb (codeTable );for _ ,_dc :=range codeTable {_cgd :=_agc ._dg .append (_dc );if _cgd !=nil {return _cgd ;};};return nil ;};type Code struct{_gbc int32 ;_bdb int32 ;_eec int32 ;_cca bool ;_ccc int32 ;};func (_bc *EncodedTable )Decode (r _b .StreamReader )(int64 ,error ){return _bc ._d .Decode (r )};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_gcc ){return nil ,_e .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_cbe :=_gcc [number -1];if _cbe ==nil {var _ebb error ;_cbe ,_ebb =_fce (_ab [number -1]);if _ebb !=nil {return nil ,_ebb ;};_gcc [number -1]=_cbe ;};return _cbe ,nil ;};func (_ebg *StandardTable )InitTree (codeTable []*Code )error {_ccb (codeTable );for _ ,_agf :=range codeTable {if _fcab :=_ebg ._dcb .append (_agf );_fcab !=nil {return _fcab ;};};return nil ;};var _ab =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};func (_ffa *FixedSizeTable )RootNode ()*InternalNode {return _ffa ._dg };var _ Node =&InternalNode {};type BasicTabler interface{HtHigh ()int32 ;HtLow ()int32 ;StreamReader ()_b .StreamReader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};func (_ce *InternalNode )Decode (r _b .StreamReader )(int64 ,error ){_cb ,_edg :=r .ReadBit ();if _edg !=nil {return 0,_edg ;};if _cb ==0{return _ce ._cdd .Decode (r );};return _ce ._gce .Decode (r );};type Tabler interface{Decode (_be _b .StreamReader )(int64 ,error );InitTree (_bded []*Code )error ;String ()string ;RootNode ()*InternalNode ;};func (_ea *ValueNode )Decode (r _b .StreamReader )(int64 ,error ){_fg ,_fgb :=r .ReadBits (byte (_ea ._ad ));if _fgb !=nil {return 0,_fgb ;};if _ea ._ca {_fg =-_fg ;};return int64 (_ea ._eda )+int64 (_fg ),nil ;};type FixedSizeTable struct{_dg *InternalNode };func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_ag :=&FixedSizeTable {_dg :&InternalNode {}};if _cfg :=_ag .InitTree (codeTable );_cfg !=nil {return nil ,_cfg ;};return _ag ,nil ;};func (_fcf *StandardTable )String ()string {return _fcf ._dcb .String ()+"\u000a"};func (_bde *FixedSizeTable )String ()string {return _bde ._dg .String ()+"\u000a"};var _ Node =&ValueNode {};func (_ff *EncodedTable )String ()string {return _ff ._d .String ()+"\u000a"};type OutOfBandNode struct{};func (_ef *InternalNode )String ()string {_ggc :=&_f .Builder {};_ggc .WriteString ("\u000a");_ef .pad (_ggc );_ggc .WriteString ("\u0030\u003a\u0020");_ggc .WriteString (_ef ._cdd .String ()+"\u000a");_ef .pad (_ggc );_ggc .WriteString ("\u0031\u003a\u0020");_ggc .WriteString (_ef ._gce .String ()+"\u000a");return _ggc .String ();};type StandardTable struct{_dcb *InternalNode };