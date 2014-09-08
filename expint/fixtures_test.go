package expint

var fixtureD = []float64{
	5.347899456798773e+01,
	5.347899456798773e+01,
	9.687303228651606e+01,
	9.687303228651606e+01,
	1.454232047228094e+01,
	1.454232047228094e+01,
	3.438572010968524e+00,
	3.438572010968524e+00,
	2.016657004886101e+00,
	2.016657004886101e+00,
	1.885277255844275e+00,
	1.885277255844275e+00,
	4.768441423047053e-01,
	4.768441423047053e-01,
	4.457790362424103e-01,
	4.457790362424103e-01,
	2.647013968214120e-01,
	2.647013968214120e-01,
	2.647013968214120e-01,
	2.647013968214120e-01,
}

var fixtureU = []float64{
	+2.876505537038424e-01,
	-2.876505537038425e-01,
	-6.456688091389395e-01,
	+6.456688091389396e-01,
	+1.920797357414034e-02,
	-1.920797357414020e-02,
	-3.983818909833339e-05,
	+3.983818909818196e-05,
	-4.803139310949334e-06,
	+4.803139311149532e-06,
	+6.731736114788428e-17,
	-1.033291828888001e-16,
	+1.374096536351588e-08,
	-1.374096515823215e-08,
	+1.586577473846631e-16,
	+1.790701711293081e-16,
	-2.562620254886738e-13,
	+2.560657865895200e-13,
	-1.079009695005139e-16,
	-1.478965236385302e-16,
	+2.871085140294832e-01,
	+2.871085140294832e-01,
	-6.459121083480398e-01,
	-6.459121083480398e-01,
	+1.913760254630729e-02,
	+1.913760254630726e-02,
	-3.964114674900070e-05,
	-3.964114674900070e-05,
	-4.787481762823613e-06,
	-4.787481762823746e-06,
	-5.967487913759480e-06,
	-5.967487913778247e-06,
	+1.369524704517120e-08,
	+1.369524704443936e-08,
	+1.707073885219112e-08,
	+1.707073878592324e-08,
	-2.554241540497770e-13,
	-2.554084873283846e-13,
	-2.888704409335241e-13,
	-2.884889260410486e-13,
	-5.857562685255296e-01,
	+5.857562685255313e-01,
	-2.518752451226868e-01,
	+2.518752451226873e-01,
	+3.053167882296536e-01,
	-3.053167882296549e-01,
	-1.502940626053376e-02,
	+1.502940626053374e-02,
	-2.052992879114804e-03,
	+2.052992879114689e-03,
	+2.982117898029278e-16,
	+4.653956354239749e-17,
	+1.255796705637273e-04,
	-1.255796705635359e-04,
	-1.492338456672307e-16,
	+7.208698461347719e-18,
	-5.117275811428179e-08,
	+5.117275783547924e-08,
	-6.983873679810908e-16,
	+1.675278057003651e-17,
	-5.941109755290541e-01,
	-5.941109755290521e-01,
	-2.556248128088455e-01,
	-2.556248128088451e-01,
	+2.854363332956238e-01,
	+2.854363332956232e-01,
	-1.426350460170470e-02,
	-1.426350460170489e-02,
	-2.047942766547712e-03,
	-2.047942766547711e-03,
	-2.553118542931820e-03,
	-2.553118542931785e-03,
	+1.303901748656634e-04,
	+1.303901748658579e-04,
	+1.625364939059144e-04,
	+1.625364939057847e-04,
	-5.595872903940702e-08,
	-5.595872868986024e-08,
	-6.321076032240504e-08,
	-6.321076121198554e-08,
	-2.705169758019947e-01,
	+2.705169758019947e-01,
	-1.392142254517067e-01,
	+1.392142254517068e-01,
	-6.282505199681553e-01,
	+6.282505199681546e-01,
	+1.101942325207011e-01,
	-1.101942325207009e-01,
	+2.413161236892554e-02,
	-2.413161236892526e-02,
	+1.077558559920273e-16,
	+2.223207376658362e-17,
	-4.018180562349694e-03,
	+4.018180562350061e-03,
	-4.870636315698015e-16,
	-3.903180915603483e-18,
	+4.815196741503325e-06,
	-4.815196742095303e-06,
	-7.750494879319825e-16,
	-6.279035163922093e-19,
	-2.493904346079419e-01,
	-2.493904346079423e-01,
	-1.294463295200206e-01,
	-1.294463295200207e-01,
	-6.271827630866398e-01,
	-6.271827630866400e-01,
	+1.342717645492044e-01,
	+1.342717645492046e-01,
	+6.020172770239295e-02,
	+6.020172770239278e-02,
	+7.519946581347704e-02,
	+7.519946581352886e-02,
	-1.221380518005515e-02,
	-1.221380518005486e-02,
	-1.524953321045432e-02,
	-1.524953321047054e-02,
	+1.936020079262857e-05,
	+1.936020079310041e-05,
	+2.190403051654730e-05,
	+2.190403051608031e-05,
	+1.182526186534761e-13,
	+1.185473110119136e-13,
	+6.205589397018488e-14,
	+6.229026578873641e-14,
	+3.243325746860037e-13,
	+3.248842167513644e-13,
	-8.050157762617971e-14,
	-8.034892196029375e-14,
	-2.514284372606528e-11,
	-2.512989816239042e-11,
	-6.881035143562134e-01,
	+6.881035144023695e-01,
	+5.953743845588805e-12,
	+5.950513776924009e-12,
	+1.628295585715988e-01,
	-1.628295585825233e-01,
	-1.186203912872941e-14,
	-1.187938636348917e-14,
	-2.972414304952820e-04,
	+2.972414305153699e-04,
	-1.687253299295838e-03,
	-1.687253299293248e-03,
	-8.876400484651934e-04,
	-8.876400484638694e-04,
	-4.657590017540530e-03,
	-4.657590017533264e-03,
	+1.167165038173957e-03,
	+1.167165038172340e-03,
	+5.140341149793060e-01,
	+5.140341149795282e-01,
	-4.573875820014794e-01,
	-4.573875819332358e-01,
	-1.217349700893069e-01,
	-1.217349700893601e-01,
	+1.082558827561021e-01,
	+1.082558827399520e-01,
	+2.453177945314841e-04,
	+2.453177945314633e-04,
	-1.976836004105587e-04,
	-1.976836003810273e-04,
	-7.590588401011105e-03,
	+7.590588401013620e-03,
	-4.040655014927785e-03,
	+4.040655014928994e-03,
	-2.231212206602658e-02,
	+2.231212206603421e-02,
	+5.581377366203838e-03,
	-5.581377366205907e-03,
	-6.876354136421230e-01,
	+6.876354136419889e-01,
	+9.888512011581232e-14,
	+8.570276649813580e-14,
	+1.629555256705528e-01,
	-1.629555256705217e-01,
	-2.277528743060259e-14,
	-2.061621868495404e-14,
	-3.287387129459821e-04,
	+3.287387129506987e-04,
	-4.145341051337717e-15,
	+4.206517921401961e-17,
	+3.202314067091629e-02,
	+3.202314067091556e-02,
	+1.685132383715750e-02,
	+1.685132383715715e-02,
	+8.855261150743018e-02,
	+8.855261150742842e-02,
	-2.224673307474358e-02,
	-2.224673307474340e-02,
	+4.533947469888070e-01,
	+4.533947469887580e-01,
	+5.085285890016312e-01,
	+5.085285890005494e-01,
	-1.077247863771825e-01,
	-1.077247863771710e-01,
	-1.207527242317751e-01,
	-1.207527242315178e-01,
	+2.182421582716948e-04,
	+2.182421582717087e-04,
	+2.216793822901741e-04,
	+2.216793822895426e-04,
	+3.013633011104263e-02,
	-3.013633011104329e-02,
	+1.653039830855156e-02,
	-1.653039830855085e-02,
	+1.058009996841112e-01,
	-1.058009996841114e-01,
	+6.979916807782847e-01,
	-6.979916807782865e-01,
	-3.070599110112638e-03,
	+3.070599110112821e-03,
	-9.905298368310731e-16,
	+5.366800753803247e-18,
	-2.056377230474006e-02,
	+2.056377230474086e-02,
	-3.255304635030873e-15,
	+5.458308761604640e-17,
	+1.107205897878383e-04,
	-1.107205897871792e-04,
	+1.710560472841682e-15,
	-2.868035640411052e-17,
	-3.686263492410183e-02,
	-3.686263492410194e-02,
	-2.017146924272114e-02,
	-2.017146924272119e-02,
	-1.292064469594519e-01,
	-1.292064469594516e-01,
	-6.905728284584387e-01,
	-6.905728284584371e-01,
	+7.771368984111982e-03,
	+7.771368984111673e-03,
	+9.656205774192302e-03,
	+9.656205774185456e-03,
	+4.199268635711441e-02,
	+4.199268635711340e-02,
	+5.226495341980501e-02,
	+5.226495341977563e-02,
	-4.563547968653436e-04,
	-4.563547968637477e-04,
	-5.146115711929661e-04,
	-5.146115711784943e-04,
	-4.400531199333174e-04,
	-4.400531199244438e-04,
	-2.428636647751982e-04,
	-2.428636647703962e-04,
	-1.615834366067326e-03,
	-1.615834366034213e-03,
	-6.851654659319383e-03,
	-6.851654659180106e-03,
	-1.270388495141432e-01,
	-1.270388495130083e-01,
	+8.054704510761429e-02,
	+8.054704512229960e-02,
	-5.357625269070853e-01,
	-5.357625269022851e-01,
	+3.402780240824624e-01,
	+3.402780241444778e-01,
	+2.369913492613943e-01,
	+2.369913492595045e-01,
	-1.353808589059048e-01,
	-1.353808589316908e-01,
	+1.122558576072907e-03,
	-1.122558576078236e-03,
	+6.265121607667195e-04,
	-6.265121607689261e-04,
	+4.324163691513533e-03,
	-4.324163691530353e-03,
	+1.849005652135374e-02,
	-1.849005652142300e-02,
	+1.490023294681157e-01,
	-1.490023294690734e-01,
	-6.043322009213550e-13,
	+1.150912698477669e-12,
	+6.282021937634307e-01,
	-6.282021937674674e-01,
	-2.552083110812786e-12,
	+4.860545213637250e-12,
	-2.877465487605142e-01,
	+2.877465487625401e-01,
	+9.938634001574740e-13,
	-2.001906011187251e-12,
	+1.572241803941817e-14,
	+2.890021650393800e-14,
	+8.662893206592360e-15,
	+1.605226196441571e-14,
	+5.746618458868369e-14,
	+1.077297973051117e-13,
	+2.460540451942883e-13,
	+4.629499908426205e-13,
	+5.353509302530313e-12,
	+7.077905969654630e-12,
	-1.508328679812238e-01,
	+1.508328679735612e-01,
	+2.258229054327796e-11,
	+2.984328441920509e-11,
	-6.369181341534484e-01,
	+6.369181341210817e-01,
	-1.014505840446489e-11,
	-1.349945155659782e-11,
	+2.675532402124672e-01,
	-2.675532401998699e-01,
	-3.644208042664913e-03,
	-3.644208042663931e-03,
	-2.011368043395216e-03,
	-2.011368043394312e-03,
	-1.338630099022903e-02,
	-1.338630099022632e-02,
	-5.675123724936089e-02,
	-5.675123724934862e-02,
	-7.268062227749815e-02,
	-7.268062227741831e-02,
	-1.194847294031560e-01,
	-1.194847294028504e-01,
	-3.055154652170889e-01,
	-3.055154652167530e-01,
	-5.032249452404752e-01,
	-5.032249452391833e-01,
	+2.025866230719996e-01,
	+2.025866230721174e-01,
	+2.990105682103532e-01,
	+2.990105682106746e-01,
	+1.229211895044157e-13,
	+1.100930808099374e-13,
	+6.789600965869706e-14,
	+6.095335731971515e-14,
	+4.531765587789671e-13,
	+4.071461050247915e-13,
	+1.920011675690669e-12,
	+1.726565288309723e-12,
	-1.281904703914760e-11,
	-1.439602603792167e-11,
	-6.134329168321384e-02,
	+6.134329171665213e-02,
	-5.445011536922215e-11,
	-6.112089562024096e-11,
	-2.604262375870872e-01,
	+2.604262377289260e-01,
	-1.433546003258568e-10,
	-1.590168205445330e-10,
	-6.545343192576006e-01,
	+6.545343195886785e-01,
	+1.995881929891205e-04,
	+1.995881930033627e-04,
	+1.101796638697213e-04,
	+1.101796638775304e-04,
	+7.338560037688235e-04,
	+7.338560038215591e-04,
	+3.105490500599830e-03,
	+3.105490500823179e-03,
	-4.343751515997350e-02,
	-4.343751515814796e-02,
	+4.472578115233897e-02,
	+4.472578111021358e-02,
	-1.843844502273918e-01,
	-1.843844502196741e-01,
	+1.897754950439534e-01,
	+1.897754948651166e-01,
	-4.674337067743635e-01,
	-4.674337067560043e-01,
	+4.556269506987255e-01,
	+4.556269502496347e-01,
	+4.930286716312571e-04,
	-4.930286716395282e-04,
	+2.752320745370893e-04,
	-2.752320745412504e-04,
	+1.901583332240941e-03,
	-1.901583332269752e-03,
	+8.119487692138497e-03,
	-8.119487692260004e-03,
	+6.602927851259291e-02,
	-6.602927851421435e-02,
	+1.202604828640528e-12,
	-2.863265180508279e-13,
	+2.799436909864286e-01,
	-2.799436909933026e-01,
	+5.106234012008937e-12,
	-1.213566636596716e-12,
	+6.459116064249566e-01,
	-6.459116064404121e-01,
	+1.317551286937225e-11,
	-2.641091455301520e-12,
	-2.136610313984631e-03,
	-2.136610313982744e-03,
	-1.179501832517189e-03,
	-1.179501832516187e-03,
	-7.856645948521368e-03,
	-7.856645948511681e-03,
	-3.321507377588168e-02,
	-3.321507377583769e-02,
	-5.661391957836644e-02,
	-5.661391957800842e-02,
	-6.062780778829802e-02,
	-6.062780778569321e-02,
	-2.393419677060489e-01,
	-2.393419677045319e-01,
	-2.563123596615259e-01,
	-2.563123596504662e-01,
	-4.293007964869280e-01,
	-4.293007964832933e-01,
	-4.297306500883081e-01,
	-4.297306500596842e-01,
}

var fixtureL = []float64{
	-3.872230919387253e+04,
	-3.870676049068540e+04,
	-1.771968893323175e+03,
	-1.682517017695009e+03,
	-6.028942742243410e+02,
	-4.559205476914974e+02,
	-3.588501346781740e+02,
	-3.587321323835815e+02,
	-3.583462839202443e+02,
	-3.568324703078450e+02,
	-1.345928057040621e+02,
	-6.694894858497091e+01,
	-2.157948007804547e+00,
	-2.102001069253590e+00,
	-2.070975359198748e+00,
	-1.614676426859295e+00,
	-2.535227433905091e-01,
	-2.413330210943396e-01,
	-2.132187943904184e-01,
	-1.236321484117528e-01,
}

var fixtureE = []float64{
	2.062601508744271e-01,
	7.841519335778077e-03,
	9.592750780197093e-02,
	3.633133924929285e-03,
	1.396405690941114e-01,
	8.619088561315427e-03,
	7.503931186978326e-03,
	3.734403227870798e-04,
	8.283218089950650e-04,
	3.171656719891085e-05,
	5.359636373895045e-04,
	5.359636373901750e-04,
	3.700767203329650e-05,
	1.118176486677200e-06,
	2.376115845121262e-05,
	2.376115845063463e-05,
	7.312358244848487e-09,
	1.752356974804581e-10,
	4.228754937256156e-09,
	4.228755071263544e-09,
	7.841519335778077e-03,
	2.062601508744271e-01,
	3.633133924929220e-03,
	9.592750780197100e-02,
	8.619088561315442e-03,
	1.396405690941116e-01,
	3.734403227868398e-04,
	7.503931186978786e-03,
	3.171656719882119e-05,
	8.283218089950484e-04,
	5.359636373895365e-04,
	5.359636373901548e-04,
	1.118176485902537e-06,
	3.700767203375247e-05,
	2.376115845164391e-05,
	2.376115845071725e-05,
	1.752358271510379e-10,
	7.312358261328360e-09,
	4.228755902521350e-09,
	4.228755023775489e-09,
	9.592750780197093e-02,
	3.633133924929227e-03,
	4.499554667023126e-02,
	1.713259163303057e-03,
	7.780101229115402e-02,
	5.035212493228518e-03,
	4.912398978247305e-03,
	2.485850093965579e-04,
	5.366111831264521e-04,
	2.025012085565576e-05,
	3.470245800474348e-04,
	3.470245800476943e-04,
	2.729398568965838e-05,
	7.879244795406951e-07,
	1.750146726107245e-05,
	1.750146726079219e-05,
	5.992364973694299e-09,
	1.329950563740134e-10,
	3.459407581902184e-09,
	3.459407646954314e-09,
	3.633133924929288e-03,
	9.592750780197100e-02,
	1.713259163303059e-03,
	4.499554667023128e-02,
	5.035212493228529e-03,
	7.780101229115399e-02,
	2.485850093970923e-04,
	4.912398978246861e-03,
	2.025012085579144e-05,
	5.366111831262556e-04,
	3.470245800473694e-04,
	3.470245800477137e-04,
	7.879244796573552e-07,
	2.729398568937855e-05,
	1.750146726099889e-05,
	1.750146726065634e-05,
	1.329953032468481e-10,
	5.992365071489335e-09,
	3.459408100801344e-09,
	3.459407837448636e-09,
	1.396405690941114e-01,
	8.619088561315456e-03,
	7.780101229115403e-02,
	5.035212493228536e-03,
	5.278330586922637e-01,
	4.385197678854264e-02,
	5.485467075833696e-02,
	3.069536607593874e-03,
	5.839326450678094e-03,
	2.322462993984567e-04,
	3.783586566807516e-03,
	3.783586566809225e-03,
	3.879740052225148e-04,
	1.125963498138779e-05,
	2.488137923583251e-04,
	2.488137923568076e-04,
	1.008237334784029e-07,
	2.191701468316071e-09,
	5.817980345754203e-08,
	5.817980372859258e-08,
	8.619088561315442e-03,
	1.396405690941115e-01,
	5.035212493228518e-03,
	7.780101229115401e-02,
	4.385197678854261e-02,
	5.278330586922635e-01,
	3.069536607594031e-03,
	5.485467075833684e-02,
	2.322462993984250e-04,
	5.839326450678177e-03,
	3.783586566807607e-03,
	3.783586566809227e-03,
	1.125963498120651e-05,
	3.879740052226441e-04,
	2.488137923587290e-04,
	2.488137923568386e-04,
	2.191701704238463e-09,
	1.008237333261809e-07,
	5.817980361149874e-08,
	5.817980370474013e-08,
	7.503931186978337e-03,
	3.734403227868398e-04,
	4.912398978247305e-03,
	2.485850093970941e-04,
	5.485467075833696e-02,
	3.069536607594031e-03,
	8.951501063498980e-01,
	2.924373895634997e-02,
	3.918834505497863e-04,
	1.136950885917160e-05,
	2.513064463530060e-04,
	2.513064463523513e-04,
	3.477477705235115e-03,
	5.598868667838937e-05,
	2.202306476756171e-03,
	2.202306476755262e-03,
	1.280633851708304e-06,
	1.371813940047317e-08,
	7.310430447070304e-07,
	7.310430451979572e-07,
	3.734403227870798e-04,
	7.503931186978793e-03,
	2.485850093965579e-04,
	4.912398978246863e-03,
	3.069536607593874e-03,
	5.485467075833683e-02,
	2.924373895634997e-02,
	8.951501063498981e-01,
	1.136950885921952e-05,
	3.918834505497175e-04,
	2.513064463530077e-04,
	2.513064463523639e-04,
	5.598868667851080e-05,
	3.477477705234860e-03,
	2.202306476756333e-03,
	2.202306476755338e-03,
	1.371813913852993e-08,
	1.280633851680549e-06,
	7.310430443357996e-07,
	7.310430451719363e-07,
	8.283218089950667e-04,
	3.171656719882287e-05,
	5.366111831264530e-04,
	2.025012085579231e-05,
	5.839326450678098e-03,
	2.322462993984249e-04,
	3.918834505497852e-04,
	1.136950885921931e-05,
	7.145580347278429e-01,
	7.814329548022636e-07,
	1.881682588176408e-05,
	1.881682588348839e-05,
	6.716574979151056e-02,
	2.851044681416448e-08,
	8.852368819332879e-07,
	8.852368820113504e-07,
	2.569197076575033e-05,
	4.374403617113387e-12,
	1.530352719547601e-10,
	1.530360352330895e-10,
	3.171656719891255e-05,
	8.283218089950484e-04,
	2.025012085565659e-05,
	5.366111831262557e-04,
	2.322462993984640e-04,
	5.839326450678177e-03,
	1.136950885917139e-05,
	3.918834505497175e-04,
	7.814329548855303e-07,
	7.145580347278432e-01,
	1.881682588180311e-05,
	1.881682588363888e-05,
	2.851044710559802e-08,
	6.716574979151113e-02,
	8.852368816314460e-07,
	8.852368817650197e-07,
	4.373473805330264e-12,
	2.569197076493154e-05,
	1.530367430002677e-10,
	1.530359797219383e-10,
	5.359636373895044e-04,
	5.359636373895348e-04,
	3.470245800474339e-04,
	3.470245800473694e-04,
	3.783586566807516e-03,
	3.783586566807600e-03,
	2.513064463530060e-04,
	2.513064463530090e-04,
	1.881682588173806e-05,
	1.881682588183303e-05,
	7.143952174190564e-01,
	2.345286288704977e-05,
	8.851956170650294e-07,
	8.851956167336972e-07,
	6.716249291389889e-02,
	1.103365253874583e-06,
	1.688715556280584e-10,
	1.688714758307786e-10,
	2.328120091654381e-05,
	1.907470362216834e-10,
	5.359636373901750e-04,
	5.359636373901548e-04,
	3.470245800476943e-04,
	3.470245800477146e-04,
	3.783586566809236e-03,
	3.783586566809227e-03,
	2.513064463523500e-04,
	2.513064463523643e-04,
	1.881682588349013e-05,
	1.881682588363845e-05,
	2.345286288705107e-05,
	7.143952174190484e-01,
	8.851956172437059e-07,
	8.851956172870740e-07,
	1.103365254223262e-06,
	6.716249291389928e-02,
	1.688714411363090e-10,
	1.688714897085664e-10,
	1.907470986717286e-10,
	2.328120091616911e-05,
	3.700767203329596e-05,
	1.118176485902320e-06,
	2.729398568965838e-05,
	7.879244796574637e-07,
	3.879740052225157e-04,
	1.125963498120586e-05,
	3.477477705235115e-03,
	5.598868667850906e-05,
	6.716574979151060e-02,
	2.851044711947581e-08,
	8.851956170650294e-07,
	8.851956172437059e-07,
	9.821934047011771e-01,
	7.048875506898700e-08,
	4.128500543085223e-06,
	4.128500545444447e-06,
	7.148883168514863e-04,
	1.284823636371613e-11,
	9.070146578249449e-10,
	9.070138251576765e-10,
	1.118176486676658e-06,
	3.700767203375203e-05,
	7.879244795407493e-07,
	2.729398568937844e-05,
	1.125963498138519e-05,
	3.879740052226458e-04,
	5.598868667838937e-05,
	3.477477705234861e-03,
	2.851044681242976e-08,
	6.716574979151116e-02,
	8.851956167284930e-07,
	8.851956172801351e-07,
	7.048875505510921e-08,
	9.821934047011777e-01,
	4.128500542412150e-06,
	4.128500545583225e-06,
	1.284863881956255e-11,
	7.148883168522080e-04,
	9.070137696465252e-10,
	9.070138529132521e-10,
	2.376115845121294e-05,
	2.376115845164381e-05,
	1.750146726107224e-05,
	1.750146726099905e-05,
	2.488137923583259e-04,
	2.488137923587317e-04,
	2.202306476756176e-03,
	2.202306476756344e-03,
	8.852368819454309e-07,
	8.852368816314460e-07,
	6.716249291389885e-02,
	1.103365254221528e-06,
	4.128500543106040e-06,
	4.128500542460722e-06,
	9.823035253720629e-01,
	5.146248073253923e-06,
	1.000903623871174e-09,
	1.000903915304718e-09,
	6.478666388072263e-04,
	1.130597296339353e-09,
	2.376115845063506e-05,
	2.376115845071790e-05,
	1.750146726079219e-05,
	1.750146726065655e-05,
	2.488137923568085e-04,
	2.488137923568399e-04,
	2.202306476755262e-03,
	2.202306476755341e-03,
	8.852368820113504e-07,
	8.852368817667544e-07,
	1.103365253891930e-06,
	6.716249291389925e-02,
	4.128500545444447e-06,
	4.128500545645675e-06,
	5.146248073253923e-06,
	9.823035253720587e-01,
	1.000903471215508e-09,
	1.000903360193206e-09,
	1.130597421239443e-09,
	6.478666388053667e-04,
	7.312358244740066e-09,
	1.752358272594581e-10,
	5.992364973694299e-09,
	1.329953032468481e-10,
	1.008237334779692e-07,
	2.191701704672144e-09,
	1.280633851704835e-06,
	1.371813913852993e-08,
	2.569197076575033e-05,
	4.373480744224167e-12,
	1.688715556280584e-10,
	1.688714411363090e-10,
	7.148883168515419e-04,
	1.284866657513817e-11,
	1.000903651626750e-09,
	1.000903471215508e-09,
	9.994743225214782e-01,
	1.165734175856414e-15,
	1.640632074639825e-13,
	1.641742297664450e-13,
	1.752356976972985e-10,
	7.312358261436780e-09,
	1.329950565908539e-10,
	5.992365071489335e-09,
	2.191701467882390e-09,
	1.008237333257472e-07,
	1.371813940047317e-08,
	1.280633851680549e-06,
	4.374393208772531e-12,
	2.569197076493501e-05,
	1.688714758307786e-10,
	1.688714966474603e-10,
	1.284826411929174e-11,
	7.148883168522635e-04,
	1.000903943060294e-09,
	1.000903360193206e-09,
	1.193489751472043e-15,
	9.994743225214773e-01,
	1.628697177125105e-13,
	1.640354518883669e-13,
	4.228754937147736e-09,
	4.228755902521350e-09,
	3.459407582119024e-09,
	3.459408100692923e-09,
	5.817980345754203e-08,
	5.817980361106506e-08,
	7.310430447052957e-07,
	7.310430443340649e-07,
	1.530352754242070e-10,
	1.530367395308208e-10,
	2.328120091654381e-05,
	1.907470847939408e-10,
	9.070146578249449e-10,
	9.070137696465252e-10,
	6.478666388072263e-04,
	1.130597490628382e-09,
	1.640354518883669e-13,
	1.628697177125105e-13,
	9.994865546329742e-01,
	1.858790898978668e-13,
	4.228755071155124e-09,
	4.228755023450229e-09,
	3.459407646954314e-09,
	3.459407837340216e-09,
	5.817980372815890e-08,
	5.817980370474013e-08,
	7.310430451927530e-07,
	7.310430451667321e-07,
	1.530360248247487e-10,
	1.530359727830444e-10,
	1.907470327522365e-10,
	2.328120091615524e-05,
	9.070137974021009e-10,
	9.070138529132521e-10,
	1.130597296339353e-09,
	6.478666388053250e-04,
	1.642297409176763e-13,
	1.640909630395981e-13,
	1.858235787466356e-13,
	9.994865546329647e-01,
}

var fixtureF = []float64{
	2.373232876136690e-02,
	3.395952652736130e-04,
	1.010152775391163e-02,
	1.451879435458541e-04,
	5.235763777408145e-03,
	1.939665345852538e-04,
	1.560489234575395e-04,
	5.629970236412176e-06,
	1.765627347222008e-05,
	4.902516466154606e-07,
	1.130891894196781e-05,
	1.130891894198992e-05,
	5.418040002633044e-07,
	1.298374418061263e-08,
	3.457607018740483e-07,
	3.457607018568179e-07,
	8.240378665889986e-11,
	1.645480747110620e-12,
	4.746837553845545e-11,
	4.746838221997983e-11,
	3.395952652736145e-04,
	2.373232876136689e-02,
	1.451879435458447e-04,
	1.010152775391163e-02,
	1.939665345852560e-04,
	5.235763777408152e-03,
	5.629970236396191e-06,
	1.560489234575672e-04,
	4.902516466109208e-07,
	1.765627347221732e-05,
	1.130891894197385e-05,
	1.130891894198925e-05,
	1.298374414398316e-08,
	5.418040002883318e-07,
	3.457607018925584e-07,
	3.457607018612571e-07,
	1.645488102449748e-12,
	8.240377899213374e-11,
	4.746840882716887e-11,
	4.746838009661403e-11,
}
