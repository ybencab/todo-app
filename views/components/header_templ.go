// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Header(authenticated bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"text-white py-5 max-w-screen-md w-full mx-auto\"><div class=\"container mx-auto flex justify-between items-center\"><a href=\"/\" class=\"text-4xl font-bold\">ToDo App</a><div class=\"flex\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if authenticated {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"rounded-md bg-gray-100 px-2 mr-2 hover:bg-violet-400\"><a href=\"/todo\" class=\"text-xl text-gray-800 font-medium\">My Board</a></div><div class=\"rounded-md bg-gray-100 px-2 hover:bg-violet-400\"><a href=\"/logout\" class=\"text-xl text-gray-800 font-medium\">Logout</a></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"rounded-md bg-gray-100 px-2 mr-2 hover:bg-violet-400\"><a href=\"/register\" class=\"text-xl text-gray-800 font-medium\">Register</a></div><div class=\"rounded-md bg-gray-100 px-2 hover:bg-violet-400\"><a href=\"/login\" class=\"text-xl text-gray-800 font-medium\">Login</a></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
