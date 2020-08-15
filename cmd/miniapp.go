package cmd

import (
	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/gcli/v2/interact"
	"github.com/gookit/gcli/v2/progress"
	"strings"
	"time"
)

func MiNiApp() *gcli.Command{
	return &gcli.Command{
		Name:        "miniapp",
		UseFor:      "生成小程序",
		Aliases:     []string{"mini"},
		Func: func (cmd *gcli.Command, args []string) error {
			ans, _ := interact.ReadLine("请输入小程序AppId：")
			if ans != "" {
				if ok := strings.HasPrefix(ans,"wx"); !ok {
					color.Info.Println("您输入的小程序AppId有误，不符合规范")
					return nil
				}
				if len(ans) != 3 {
					color.Info.Println("您输入的小程序AppId有误，不符合规范")
					return nil
				}
				color.Println("Your input: ", ans)
				ans1 := interact.MultiSelect(
					"请选择您要使用的模板",
					[]string{"Cell", "Tab", "Button"},
					nil,
					false,
				)
				if len(ans1) <=0 {
					color.Info.Println("您未选择模板")
					return nil
				}
				color.Comment.Printf("您选择的模板是: %s,正在构建中,请等待", strings.Join(ans1,"|"))
				time.Sleep(1000 * time.Millisecond)
				speed := 100
				maxSteps := 110
				p := progress.Bar(maxSteps)
				p.Start()

				for i := 0; i < maxSteps; i++ {
					time.Sleep(time.Duration(speed) * time.Millisecond)
					p.Advance()
				}

				p.Finish()
			} else {
				color.Cyan.Println("抱歉，您未输入小程序AppId")
			}
			return nil
		},
	}
}
