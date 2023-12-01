package parser

// instance.Instances[].ImageId[1]

type JqLite struct {
	Expr []SelectorExpr `@@ ( "." @@ )*`
}

type SelectorExpr struct {
	Sel string `@Ident`
	X []IndexExpr  `("[" (@@)? "]")*`
}

type IndexExpr struct {
	Index *int    `@(Int)?`
}
