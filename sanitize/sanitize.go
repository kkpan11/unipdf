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

package sanitize ;import (_d "github.com/unidoc/unipdf/v3/common";_ba "github.com/unidoc/unipdf/v3/core";);func (_dba *Sanitizer )analyze (_ddd []_ba .PdfObject ){_dbc :=map[string ]int {};for _ ,_aaf :=range _ddd {switch _edb :=_aaf .(type ){case *_ba .PdfIndirectObject :_aag ,_bg :=_ba .GetDict (_edb .PdfObject );
if _bg {if _ca ,_ge :=_ba .GetName (_aag .Get ("\u0054\u0079\u0070\u0065"));_ge &&*_ca =="\u0043a\u0074\u0061\u006c\u006f\u0067"{if _ ,_ad :=_ba .GetIndirect (_aag .Get ("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e"));_ad {_dbc ["\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e"]++;
};}else if _ebb ,_bb :=_ba .GetName (_aag .Get ("\u0053"));_bb {_gee :=_ebb .String ();if _gee =="\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"||_gee =="\u0055\u0052\u0049"||_gee =="\u0047\u006f\u0054\u006f"||_gee =="\u0047\u006f\u0054o\u0052"||_gee =="\u004c\u0061\u0075\u006e\u0063\u0068"{_dbc [_gee ]++;
}else if _gee =="\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn"{if _ ,_cg :=_ba .GetStream (_aag .Get ("\u004a\u0053"));_cg {_dbc [_gee ]++;};};}else if _adf :=_aag .Get ("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074");_adf !=nil {_dbc ["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++;
}else if _eg ,_cgd :=_ba .GetIndirect (_aag .Get ("\u0050\u0061\u0072\u0065\u006e\u0074"));_cgd {if _fdg ,_fc :=_ba .GetDict (_eg .PdfObject );_fc {if _da ,_deb :=_ba .GetDict (_fdg .Get ("\u0041\u0041"));_deb {_gab :=_da .Get ("\u004b");_aff ,_afd :=_ba .GetIndirect (_gab );
if _afd {if _egc ,_aefa :=_ba .GetDict (_aff .PdfObject );_aefa {if _ff ,_cae :=_ba .GetName (_egc .Get ("\u0053"));_cae &&*_ff =="\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"{_dbc ["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++;
}else if _ ,_ecf :=_ba .GetString (_egc .Get ("\u004a\u0053"));_ecf {_dbc ["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++;}else {_bf :=_da .Get ("\u0046");if _bf !=nil {_gafg ,_caea :=_ba .GetIndirect (_bf );if _caea {if _cc ,_gg :=_ba .GetDict (_gafg .PdfObject );
_gg {if _eec ,_gc :=_ba .GetName (_cc .Get ("\u0053"));_gc {_cba :=_eec .String ();_dbc [_cba ]++;};};};};};};};};};};};};};_dba ._c =_dbc ;};

// Sanitizer represents a sanitizer object.
// It implements the Optimizer interface to access the objects field from the writer.
type Sanitizer struct{_de SanitizationOpts ;_c map[string ]int ;};

// New returns a new sanitizer object.
func New (opts SanitizationOpts )*Sanitizer {return &Sanitizer {_de :opts }};

// SanitizationOpts specifies the objects to be removed during sanitization.
type SanitizationOpts struct{

// JavaScript specifies wether JavaScript action should be removed. JavaScript Actions, section 12.6.4.16 of PDF32000_2008
JavaScript bool ;

// URI specifies if URI actions should be removed. 12.6.4.7 URI Actions, PDF32000_2008.
URI bool ;

// GoToR removes remote GoTo actions. 12.6.4.3 Remote Go-To Actions, PDF32000_2008.
GoToR bool ;

// GoTo specifies wether GoTo actions should be removed. 12.6.4.2 Go-To Actions, PDF32000_2008.
GoTo bool ;

// RenditionJS enables removing of `JS` entry from a Rendition Action.
// The `JS` entry has a value of text string or stream containing a JavaScript script that shall be executed when the action is triggered.
// 12.6.4.13 Rendition Actions Table 214, PDF32000_2008.
RenditionJS bool ;

// OpenAction removes OpenAction entry from the document catalog.
OpenAction bool ;

// Launch specifies wether Launch Action should be removed.
// A launch action launches an application or opens or prints a document.
// 12.6.4.5 Launch Actions, PDF32000_2008.
Launch bool ;};func (_ee *Sanitizer )processObjects (_f []_ba .PdfObject )([]_ba .PdfObject ,error ){_dc :=[]_ba .PdfObject {};_dd :=_ee ._de ;for _ ,_fa :=range _f {switch _be :=_fa .(type ){case *_ba .PdfIndirectObject :_ecb ,_ea :=_ba .GetDict (_be );
if _ea {if _eaf ,_bc :=_ba .GetName (_ecb .Get ("\u0054\u0079\u0070\u0065"));_bc &&*_eaf =="\u0043a\u0074\u0061\u006c\u006f\u0067"{if _ ,_ed :=_ba .GetIndirect (_ecb .Get ("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e"));_ed &&_dd .OpenAction {_ecb .Remove ("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e");
};}else if _fd ,_a :=_ba .GetName (_ecb .Get ("\u0053"));_a {switch *_fd {case "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074":if _dd .JavaScript {if _cd ,_aa :=_ba .GetStream (_ecb .Get ("\u004a\u0053"));_aa {_ef :=[]byte {};_eag ,_df :=_ba .MakeStream (_ef ,nil );
if _df ==nil {*_cd =*_eag ;};};_d .Log .Debug ("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e");continue ;};case "\u0055\u0052\u0049":if _dd .URI {_d .Log .Debug ("\u0055\u0052\u0049\u0020ac\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070\u0070\u0065\u0064\u002e");
continue ;};case "\u0047\u006f\u0054\u006f":if _dd .GoTo {_d .Log .Debug ("G\u004fT\u004f\u0020\u0061\u0063\u0074\u0069\u006f\u006e \u0073\u006b\u0069\u0070pe\u0064\u002e");continue ;};case "\u0047\u006f\u0054o\u0052":if _dd .GoToR {_d .Log .Debug ("R\u0065\u006d\u006f\u0074\u0065\u0020G\u006f\u0054\u004f\u0020\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u006bi\u0070p\u0065\u0064\u002e");
continue ;};case "\u004c\u0061\u0075\u006e\u0063\u0068":if _dd .Launch {_d .Log .Debug ("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e");continue ;};case "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn":if _fb ,_ddc :=_ba .GetStream (_ecb .Get ("\u004a\u0053"));
_ddc {_cb :=[]byte {};_dea ,_eaa :=_ba .MakeStream (_cb ,nil );if _eaa ==nil {*_fb =*_dea ;};};};}else if _ag :=_ecb .Get ("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074");_ag !=nil &&_dd .JavaScript {continue ;}else if _fg ,_efc :=_ba .GetName (_ecb .Get ("\u0054\u0079\u0070\u0065"));
_efc &&*_fg =="\u0041\u006e\u006eo\u0074"&&_dd .JavaScript {if _ae ,_def :=_ba .GetIndirect (_ecb .Get ("\u0050\u0061\u0072\u0065\u006e\u0074"));_def {if _af ,_eef :=_ba .GetDict (_ae .PdfObject );_eef {if _db ,_g :=_ba .GetDict (_af .Get ("\u0041\u0041"));
_g {_ce ,_aef :=_ba .GetIndirect (_db .Get ("\u004b"));if _aef {if _gf ,_ga :=_ba .GetDict (_ce .PdfObject );_ga {if _eefb ,_fdc :=_ba .GetName (_gf .Get ("\u0053"));_fdc &&*_eefb =="\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"{_gf .Clear ();
}else if _dbb :=_db .Get ("\u0046");_dbb !=nil {if _cdb ,_cda :=_ba .GetIndirect (_dbb );_cda {if _gaf ,_efce :=_ba .GetDict (_cdb .PdfObject );_efce {if _bcb ,_cec :=_ba .GetName (_gaf .Get ("\u0053"));_cec &&*_bcb =="\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"{_gaf .Clear ();
};};};};};};};};};};};case *_ba .PdfObjectStream :_d .Log .Debug ("\u0070d\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073t\u0072e\u0061m\u0020\u0074\u0079\u0070\u0065\u0020\u0025T",_be );case *_ba .PdfObjectStreams :_d .Log .Debug ("\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0074\u0079\u0070e\u0020\u0025\u0054",_be );
default:_d .Log .Debug ("u\u006e\u006b\u006e\u006fwn\u0020p\u0064\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0025\u0054",_be );};_dc =append (_dc ,_fa );};_ee .analyze (_dc );return _dc ,nil ;};

// GetSuspiciousObjects returns a count of each detected suspicious object.
func (_eeg *Sanitizer )GetSuspiciousObjects ()map[string ]int {return _eeg ._c };

// Optimize optimizes `objects` and returns updated list of objects.
func (_e *Sanitizer )Optimize (objects []_ba .PdfObject )([]_ba .PdfObject ,error ){return _e .processObjects (objects );};