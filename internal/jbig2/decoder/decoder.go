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

package decoder ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_e "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_ee "github.com/unidoc/unipdf/v3/internal/jbig2/document";_g "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_c "image";
);func (_gga *Decoder )DecodeNextPage ()([]byte ,error ){_gga ._de ++;_aa :=_gga ._de ;return _gga .decodePage (_aa );};func (_ge *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _ge .decodePage (pageNumber )};func (_fd *Decoder )DecodePageImage (pageNumber int )(_c .Image ,error ){const _a ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
_da ,_ad :=_fd .decodePageImage (pageNumber );if _ad !=nil {return nil ,_g .Wrap (_ad ,_a ,"");};return _da ,nil ;};func (_ea *Decoder )PageNumber ()(int ,error ){const _fg ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";
if _ea ._gg ==nil {return 0,_g .Error (_fg ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");};return int (_ea ._gg .NumberOfPages ),nil ;};type Decoder struct{_ff *_f .Reader ;
_gg *_ee .Document ;_de int ;_cg Parameters ;};func (_fdd *Decoder )decodePage (_ef int )([]byte ,error ){const _efg ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _ef < 0{return nil ,_g .Errorf (_efg ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ef );
};if _ef > int (_fdd ._gg .NumberOfPages ){return nil ,_g .Errorf (_efg ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ef );
};_b ,_ae :=_fdd ._gg .GetPage (_ef );if _ae !=nil {return nil ,_g .Wrap (_ae ,_efg ,"");};_ga ,_ae :=_b .GetBitmap ();if _ae !=nil {return nil ,_g .Wrap (_ae ,_efg ,"");};_ga .InverseData ();if !_fdd ._cg .UnpaddedData {return _ga .Data ,nil ;};return _ga .GetUnpaddedData ();
};type Parameters struct{UnpaddedData bool ;Color _e .Color ;};func Decode (input []byte ,parameters Parameters ,globals *_ee .Globals )(*Decoder ,error ){_fc :=_f .NewReader (input );_df ,_eef :=_ee .DecodeDocument (_fc ,globals );if _eef !=nil {return nil ,_eef ;
};return &Decoder {_ff :_fc ,_gg :_df ,_cg :parameters },nil ;};func (_ca *Decoder )decodePageImage (_fdb int )(_c .Image ,error ){const _gc ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _fdb < 0{return nil ,_g .Errorf (_gc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_fdb );
};if _fdb > int (_ca ._gg .NumberOfPages ){return nil ,_g .Errorf (_gc ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_fdb );
};_ed ,_dad :=_ca ._gg .GetPage (_fdb );if _dad !=nil {return nil ,_g .Wrap (_dad ,_gc ,"");};_cc ,_dad :=_ed .GetBitmap ();if _dad !=nil {return nil ,_g .Wrap (_dad ,_gc ,"");};_cc .InverseData ();return _cc .ToImage (),nil ;};