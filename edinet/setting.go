package edinet

var weights = map[string]map[string]float64{
	// アクルーアルを計算するための設定。
	// 通常は アクルーアル＝純利益－営業キャッシュフロー
	// ただしここでは財テクや特別利益・損失を除くため、
	// 純利益の代わりに、経常利益×（1-法人税率）を使用する。
	"Accruals": {
		"OrdinaryIncome":                             0.6, // 経常利益×（1-法人税率）
		"NetCashProvidedByUsedInOperatingActivities": -1,
	},
	// 解散価値を計算するための設定。
	"BreakupValue": {
		// 流動資産
		// 流動資産（合計）30%
		"CurrentAssets": 0.3,
		// 流動資産の例外について差分を上乗せ。現金預金・有価証券 100%,
		"CashAndDeposits":               0.7,
		"ShortTermInvestmentSecurities": 0.7,
		// 手形類 90%,
		"AccountsReceivableTrade":                                                        0.6,
		"NotesAndAccountsReceivableTrade":                                                0.6,
		"NotesAndOperatingAccountsPayableTrade":                                          0.6,
		"NotesReceivableAccountsReceivableFromCompletedConstructionContractsAndOtherCNS": 0.6,

		// 固定資産
		// 有形固定資産（合計）20%
		"PropertyPlantAndEquipment": 0.2,
		// 例外の差分を上乗せ
		// 土地・賃貸不動産（純額） 50%
		"Land":                 0.3,
		"RealEstateForRentNet": 0.3,
		// 無形固定資産は一切考慮しない。曖昧なものは全てゼロ

		// 投資その他の資産
		// 投資有価証券・長期預金 100%
		"InvestmentSecurities":                                    1,
		"StocksOfSubsidiariesAndAffiliates":                       1,
		"BondsOfSubsidiariesAndAffiliates":                        1,
		"InvestmentsInOtherSecuritiesOfSubsidiariesAndAffiliates": 1,
		"OperationalInvestmentSecuritiesIOA":                      1,
		"LongTermDeposits":                                        1,
		"LongTermTimeDeposits":                                    1,
		"LongTermLoansReceivable":                                 1,
		"LongTermLoansReceivableFromSubsidiariesAndAffiliates":    1,
		// 敷金及び保証金・建設協力金 30%
		"LeaseAndGuaranteeDeposits":             0.3,
		"ConstructionAssistanceFundReceivables": 0.3,

		// 負債は耳を揃えて返さなければいけないので全部の合計を引く
		"Liabilities": -1,
	},
	// ネットキャッシュを計算するための設定
	// ネットキャッシュ＝現金及び現金同等物＋有価証券－有利子負債
	"NetCash": {
		// 現金及び現金同等物
		"CashAndDeposits": 1,
		// 有価証券
		"ShortTermInvestmentSecurities":                           1,
		"InvestmentSecurities":                                    1,
		"StocksOfSubsidiariesAndAffiliates":                       1,
		"BondsOfSubsidiariesAndAffiliates":                        1,
		"InvestmentsInOtherSecuritiesOfSubsidiariesAndAffiliates": 1,
		"OperationalInvestmentSecuritiesIOA":                      1,
		"LongTermDeposits":                                        1,
		"LongTermTimeDeposits":                                    1,
		// 有利子負債＝短期借入金・長期借入金・社債・コマーシャルペーパー等
		"ShortTermBondsPayable":                                           -1,
		"ShortTermLoansPayable":                                           -1,
		"ShortTermLoansPayableToSubsidiariesAndAffiliates":                -1,
		"CommercialPapersLiabilities":                                     -1,
		"CurrentPortionOfBonds":                                           -1,
		"CurrentPortionOfLongTermLoansPayable":                            -1,
		"CurrentPortionOfLongTermLoansPayableToSubsidiariesAndAffiliates": -1,
		"CurrentPortionOfConvertibleBonds":                                -1,
		"CurrentPortionOfBondsWithSubscriptionRightsToShares":             -1,
		"CurrentPortionOfOtherNoncurrentLiabilities":                      -1,
		"BondsPayable":                                                    -1,
		"ConvertibleBonds":                                                -1,
		"ConvertibleBondTypeBondsWithSubscriptionRightsToShares":          -1,
		"BondsWithSubscriptionRightsToSharesNCL":                          -1,
		"LongTermLoansPayable":                                            -1,
		"LongTermLoansPayableToShareholdersDirectorsOrEmployees":          -1,
		"LongTermLoansPayableToSubsidiariesAndAffiliates":                 -1,
	},
	// フリーキャッシュフローを計算するためのアイテム
	// 通常は フリーキャッシュフロー＝営業キャッシュフロー＋投資キャッシュフロー
	// ただしここでは財テクを除くため、
	// 営業キャッシュフロー＋有形及び無形固定資産の取得による支出
	"FreeCashFlow": {
		// 営業キャッシュフロー
		"NetCashProvidedByUsedInOperatingActivities": 1,
		// "NetCashProvidedByUsedInOperatingActivitiesSummaryOfBusinessResults": 1,

		// 有形及び無形固定資産の取得による支出
		"PurchaseOfPropertyPlantAndEquipmentAndIntangibleAssetsInvCF": 1,
		"PurchaseOfPropertyPlantAndEquipmentInvCF":                    1,
		"PurchaseOfNoncurrentAssetsInvCF":                             1,
		"PurchaseOfIntangibleAssetsInvCF":                             1,
	},
}
