// +build debug

package polyclip

// Enable with: [go test -tags debug] for [go test -tags "debug footag bartag"]
func _DBG(f func()) {
	f()
}
