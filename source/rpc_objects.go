// rpc_objects.go
package main

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
