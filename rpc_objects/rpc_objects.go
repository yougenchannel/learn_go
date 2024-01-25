package rpc_objects

type Args struct {
	M, N int
}

/*
	Rpc要求方法需要返回error
*/
func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
