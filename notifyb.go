// notifyb パッケージは通知を表す error を提供します.
//
// エラーは通常 ErrXXX のような名称で定義しますが、
// io.EOF などのように通知を表す error を定義したい場合があります.
//
// 次のように記述すると MyEOF は error として振る舞うことができます.
//
//	const MyEOF notify.Notify = "my eof"
//
//	func Foo() error {
//		return MyEOF
//	}
package notifyb

// Notify は通知を error で返すための型です.
type Notify string

func (e Notify) Error() string {
	return string(e)
}
