package hoi

type Options struct {
	Server bool `short:"s" long:"server" description:"Start hoi server."`
	Clear  bool `short:"c" long:"clear" description:"Clear all symlinks by removing public directory."`
}
