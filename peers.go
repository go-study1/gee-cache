package geecache

import pb "github.com/go-study1/gee-cache/geecachepb"

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}
