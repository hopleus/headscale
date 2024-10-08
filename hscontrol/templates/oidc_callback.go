package templates

import (
	"fmt"
	"strings"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
	"github.com/juanfont/headscale/hscontrol/types"
)

type OIDCCallbackConfig struct {
	Verb string
	User *types.User
	Node *types.Node
}

var oidcStyleMgr = styles.NewStyleManager()

var oidcBodyTagClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.Color:                          "#242424",
	styles.BackgroundColor:                "white",
	styles.FontFamily:                     "system-ui,-apple-system,BlinkMacSystemFont,\"Segoe UI\",\"Roboto\",\"Oxygen\",\"Ubuntu\",\"Cantarell\",\"Fira Sans\",\"Droid Sans\",\"Helvetica Neue\",sans-serif",
	styles.LetterSpacing:                  "-0.0015em",
	styles.LineHeight:                     "1.45",
	styles.Var("-webkit-font-smoothing"):  "antialiased",
	styles.Var("-moz-osx-font-smoothing"): "grayscale",
})

var oidcMainTagClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.Display:       "flex",
	styles.FlexDirection: "column",
	styles.AlignItems:    "center",
	styles.PaddingTop:    "2rem",
	styles.PaddingBottom: "1rem",
	styles.TextAlign:     "center",
	styles.MinHeight:     "100%",
	styles.Margin:        "0 auto",
})

var oidcContainerStyle = styles.Props{
	styles.MarginLeft:   "auto",
	styles.MarginRight:  "auto",
	styles.PaddingLeft:  "1rem",
	styles.PaddingRight: "1rem",
	styles.MaxWidth:     "32rem",
	styles.Width:        "100%",
}

var oidcHeaderBlockClass = oidcStyleMgr.AddStyle(
	styles.Merge(
		styles.Props{
			styles.Display:        "flex",
			styles.AlignItems:     "center",
			styles.JustifyContent: "space-between",
			styles.Gap:            "1rem",
			styles.MarginBottom:   "4rem",
		},
		oidcContainerStyle,
	),
)

var oidcHeaderUserBlockClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.Display:    "flex",
	styles.AlignItems: "center",
	styles.FontSize:   "0.85rem",
	styles.Gap:        "0.5rem",
	styles.MinWidth:   "0",
})

var oidcHeaderUserBlockTruncateClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.Overflow:     "hidden",
	styles.WhiteSpace:   "nowrap",
	styles.TextOverflow: "ellipsis",
})

var oidcHeaderUserBlockPictureClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.FlexShrink:      "0",
	styles.Width:           "1.5rem",
	styles.Height:          "1.5rem",
	styles.BackgroundSize:  "cover",
	styles.BackgroundColor: "#e5e5e5",
	styles.Border:          "1px #eeebea solid",
	styles.BorderRadius:    "999px",
})

var (
	oidcContainerClass = oidcStyleMgr.AddStyle(oidcContainerStyle)
	oidcH1Class        = oidcStyleMgr.AddStyle(styles.Props{
		styles.FontSize:      "1.25rem",
		styles.FontWeight:    "600",
		styles.LetterSpacing: "-0.03em",
	})
)

var oidcHRClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.Margin:          "2rem 0",
	styles.Border:          "0",
	styles.Height:          "1px",
	styles.BackgroundColor: "#e8e8e8",
})

var oidcAClass = oidcStyleMgr.AddCompositeStyle(styles.CompositeStyle{
	Default: styles.Props{
		styles.Display:        "flex",
		styles.Margin:         "8px 0",
		styles.Color:          "#1563ff",
		styles.TextDecoration: "none",
	},
	PseudoClasses: map[string]styles.Props{
		"hover": {
			styles.Color: "black",
		},
	},
})

var oidcIconClass = oidcStyleMgr.AddStyle(styles.Props{
	styles.Display:        "flex",
	styles.AlignItems:     "center",
	styles.JustifyContent: "center",
	styles.Width:          "21px",
	styles.Height:         "21px",
	styles.VerticalAlign:  "middle",
})

func OIDCCallback(cfg OIDCCallbackConfig) string {
	logo := &elem.Element{
		Tag: "svg",
		Attrs: attrs.Props{
			attrs.Width:  "146",
			attrs.Height: "51",
			"xmlns":      "http://www.w3.org/2000/svg",
			"xml:space":  "preserve",
			attrs.Style: styles.Props{
				styles.Var("fill-rule"):         "evenodd",
				styles.Var("clip-rule"):         "evenodd",
				styles.Var("stroke-linejoin"):   "round",
				styles.Var("stroke-miterlimit"): styles.Int(2),
			}.ToInline(),
			"viewBox": "0 0 1280 640",
		},
		Children: []elem.Node{
			elem.Raw("<path d=\"M.08 0v-.736h.068v.3C.203-.509.27-.545.347-.545c.029 0 .055.005.079.015.024.01.045.025.062.045.017.02.031.045.041.075.009.03.014.065.014.105V0H.475v-.289C.475-.352.464-.4.443-.433.422-.466.385-.483.334-.483c-.027 0-.052.006-.075.017C.236-.455.216-.439.2-.419c-.017.02-.029.044-.038.072-.009.028-.014.059-.014.093V0H.08Z\" style=\"fill: #f8b5cb; fill-rule: nonzero\" transform=\"translate(32.92220721 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.051-.264c0-.036.007-.071.02-.105.013-.034.031-.064.055-.09.023-.026.052-.047.086-.063.033-.015.071-.023.112-.023.039 0 .076.007.109.021.033.014.062.033.087.058.025.025.044.054.058.088.014.035.021.072.021.113v.005H.121c.001.031.007.059.018.084.01.025.024.047.042.065.018.019.04.033.065.043.025.01.052.015.082.015.026 0 .049-.003.069-.01.02-.007.038-.016.054-.028C.466-.102.48-.115.492-.13c.011-.015.022-.03.032-.046l.057.03C.556-.097.522-.058.48-.03.437-.001.387.013.328.013.284.013.245.006.21-.01.175-.024.146-.045.123-.07.1-.095.082-.125.07-.159.057-.192.051-.227.051-.264ZM.128-.32h.396C.51-.375.485-.416.449-.441.412-.466.371-.479.325-.479c-.048 0-.089.013-.123.039-.034.026-.059.066-.074.12Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(177.16674681 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.051-.267c0-.038.007-.074.021-.108.014-.033.033-.063.058-.088.025-.025.054-.045.087-.06.033-.015.069-.022.108-.022.043 0 .083.009.119.027.035.019.066.047.093.084v-.097h.067V0H.537v-.091C.508-.056.475-.029.44-.013.404.005.365.013.323.013.284.013.248.006.215-.01.182-.024.153-.045.129-.071.104-.096.085-.126.072-.16.058-.193.051-.229.051-.267Zm.279.218c.027 0 .054-.005.079-.015.025-.01.048-.024.068-.043.019-.018.035-.04.047-.067.012-.027.018-.056.018-.089 0-.031-.005-.059-.016-.086C.515-.375.501-.398.482-.417.462-.436.44-.452.415-.463.389-.474.361-.479.331-.479c-.031 0-.059.006-.084.017C.221-.45.199-.434.18-.415c-.019.02-.033.043-.043.068-.011.026-.016.053-.016.082 0 .029.005.056.016.082.011.026.025.049.044.069.019.02.041.036.066.047.025.012.053.018.083.018Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(327.76463481 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.051-.267c0-.038.007-.074.021-.108.014-.033.033-.063.058-.088.025-.025.054-.045.087-.06.033-.015.069-.022.108-.022.043 0 .083.009.119.027.035.019.066.047.093.084v-.302h.068V0H.537v-.091C.508-.056.475-.029.44-.013.404.005.365.013.323.013.284.013.248.006.215-.01.182-.024.153-.045.129-.071.104-.096.085-.126.072-.16.058-.193.051-.229.051-.267Zm.279.218c.027 0 .054-.005.079-.015.025-.01.048-.024.068-.043.019-.018.035-.04.047-.067.011-.027.017-.056.017-.089 0-.031-.005-.059-.016-.086C.514-.375.5-.398.481-.417.462-.436.439-.452.414-.463.389-.474.361-.479.331-.479c-.031 0-.059.006-.084.017C.221-.45.199-.434.18-.415c-.019.02-.033.043-.043.068-.011.026-.016.053-.016.082 0 .029.005.056.016.082.011.026.025.049.044.069.019.02.041.036.066.047.025.012.053.018.083.018Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(488.71612761 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"m.034-.062.043-.049c.017.019.035.034.054.044.018.01.037.015.057.015.013 0 .026-.002.038-.007.011-.004.021-.01.031-.018.009-.008.016-.017.021-.028.005-.011.008-.022.008-.035 0-.019-.005-.034-.014-.047C.263-.199.248-.21.229-.221.205-.234.183-.247.162-.259.14-.271.122-.284.107-.298.092-.311.08-.327.071-.344.062-.361.058-.381.058-.404c0-.021.004-.04.012-.058.007-.016.018-.031.031-.044.013-.013.028-.022.046-.029.018-.007.037-.01.057-.01.029 0 .056.006.079.019s.045.031.068.053l-.044.045C.291-.443.275-.456.258-.465.241-.474.221-.479.2-.479c-.022 0-.041.007-.056.02C.128-.445.12-.428.12-.408c0 .019.006.035.017.048.011.013.027.026.048.037.027.015.05.028.071.04.021.013.038.026.052.039.014.013.025.028.032.044.007.016.011.035.011.057 0 .021-.004.041-.011.059-.008.019-.019.036-.033.05-.014.015-.031.026-.05.035C.237.01.215.014.191.014c-.03 0-.059-.006-.086-.02C.077-.019.053-.037.034-.062Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(649.90292961 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.051-.266c0-.04.007-.077.022-.111.014-.034.034-.063.059-.089.025-.025.054-.044.089-.058.035-.014.072-.021.113-.021.051 0 .098.01.139.03.041.021.075.049.1.085l-.05.043C.498-.418.47-.441.439-.456.408-.471.372-.479.331-.479c-.03 0-.058.005-.083.016C.222-.452.2-.436.181-.418.162-.399.148-.376.137-.35c-.011.026-.016.054-.016.084 0 .031.005.06.016.086.011.027.025.049.044.068.019.019.041.034.067.044.025.011.053.016.084.016.077 0 .141-.03.191-.09l.051.04c-.028.036-.062.064-.103.085C.43.004.384.014.332.014.291.014.254.007.219-.008.184-.022.155-.042.13-.067.105-.092.086-.121.072-.156.058-.19.051-.227.051-.266Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(741.20289921 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.051-.267c0-.038.007-.074.021-.108.014-.033.033-.063.058-.088.025-.025.054-.045.087-.06.033-.015.069-.022.108-.022.043 0 .083.009.119.027.035.019.066.047.093.084v-.097h.067V0H.537v-.091C.508-.056.475-.029.44-.013.404.005.365.013.323.013.284.013.248.006.215-.01.182-.024.153-.045.129-.071.104-.096.085-.126.072-.16.058-.193.051-.229.051-.267Zm.279.218c.027 0 .054-.005.079-.015.025-.01.048-.024.068-.043.019-.018.035-.04.047-.067.012-.027.018-.056.018-.089 0-.031-.005-.059-.016-.086C.515-.375.501-.398.482-.417.462-.436.44-.452.415-.463.389-.474.361-.479.331-.479c-.031 0-.059.006-.084.017C.221-.45.199-.434.18-.415c-.019.02-.033.043-.043.068-.011.026-.016.053-.016.082 0 .029.005.056.016.082.011.026.025.049.044.069.019.02.041.036.066.047.025.012.053.018.083.018Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(884.27089281 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.066-.736h.068V0H.066z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(1045.22238561 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<path d=\"M.051-.264c0-.036.007-.071.02-.105.013-.034.031-.064.055-.09.023-.026.052-.047.086-.063.033-.015.071-.023.112-.023.039 0 .076.007.109.021.033.014.062.033.087.058.025.025.044.054.058.088.014.035.021.072.021.113v.005H.121c.001.031.007.059.018.084.01.025.024.047.042.065.018.019.04.033.065.043.025.01.052.015.082.015.026 0 .049-.003.069-.01.02-.007.038-.016.054-.028C.466-.102.48-.115.492-.13c.011-.015.022-.03.032-.046l.057.03C.556-.097.522-.058.48-.03.437-.001.387.013.328.013.284.013.245.006.21-.01.175-.024.146-.045.123-.07.1-.095.082-.125.07-.159.057-.192.051-.227.051-.264ZM.128-.32h.396C.51-.375.485-.416.449-.441.412-.466.371-.479.325-.479c-.048 0-.089.013-.123.039-.034.026-.059.066-.074.12Z\" style=\"fill: #8d8d8d; fill-rule: nonzero\" transform=\"translate(1092.28422561 521.8022953) scale(235.3092)\" />"),
			elem.Raw("<circle cx=\"141.023\" cy=\"338.36\" r=\"117.472\" style=\"fill: #f8b5cb\" transform=\"matrix(.581302 0 0 .58613 40.06479894 12.59842153)\" />"),
			elem.Raw("<circle cx=\"352.014\" cy=\"268.302\" r=\"33.095\" style=\"fill: #a2a2a2\" transform=\"matrix(.59308 0 0 .58289 32.39345942 21.2386)\" />"),
			elem.Raw("<circle cx=\"352.014\" cy=\"268.302\" r=\"33.095\" style=\"fill: #a2a2a2\" transform=\"matrix(.59308 0 0 .58289 32.39345942 88.80371146)\" />"),
			elem.Raw("<circle cx=\"352.014\" cy=\"268.302\" r=\"33.095\" style=\"fill: #a2a2a2\" transform=\"matrix(.59308 0 0 .58289 120.7528627 88.80371146)\" />"),
			elem.Raw("<circle cx=\"352.014\" cy=\"268.302\" r=\"33.095\" style=\"fill: #a2a2a2\" transform=\"matrix(.59308 0 0 .58289 120.99825939 21.2386)\" />"),
			elem.Raw("<circle cx=\"805.557\" cy=\"336.915\" r=\"118.199\" style=\"fill: #8d8d8d\" transform=\"matrix(.5782 0 0 .58289 36.19871106 15.26642564)\" />"),
			elem.Raw("<circle cx=\"805.557\" cy=\"336.915\" r=\"118.199\" style=\"fill: #8d8d8d\" transform=\"matrix(.5782 0 0 .58289 183.24041937 15.26642564)\" />"),
			elem.Raw("<path d=\"M680.282 124.808h-68.093v390.325h68.081v-28.23H640V153.228h40.282v-28.42Z\" style=\"fill: #303030\" transform=\"translate(34.2345 21.2386) scale(.58289)\" />"),
			elem.Raw("<path d=\"M680.282 124.808h-68.093v390.325h68.081v-28.23H640V153.228h40.282v-28.42Z\" style=\"fill: #303030\" transform=\"matrix(-.58289 0 0 .58289 1116.7719791 21.2386)\" />"),
		},
	}
	iconSuccess := &elem.Element{
		Tag: "svg",
		Attrs: attrs.Props{
			attrs.Width:  "30",
			attrs.Height: "30",
			"viewBox":    "0 0 24 24",
			"fill":       "none",
			"xmlns":      "http://www.w3.org/2000/svg",
			attrs.Style: styles.Props{
				styles.MarginBottom: "1rem",
			}.ToInline(),
		},
		Children: []elem.Node{
			elem.Raw("<path d=\"M22 11.0799V11.9999C21.9988 14.1563 21.3005 16.2545 20.0093 17.9817C18.7182 19.7088 16.9033 20.9723 14.8354 21.5838C12.7674 22.1952 10.5573 22.1218 8.53447 21.3744C6.51168 20.6271 4.78465 19.246 3.61096 17.4369C2.43727 15.6279 1.87979 13.4879 2.02168 11.3362C2.16356 9.18443 2.99721 7.13619 4.39828 5.49694C5.79935 3.85768 7.69279 2.71525 9.79619 2.24001C11.8996 1.76477 14.1003 1.9822 16.07 2.85986\" stroke=\"#33C27F\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" />"),
			elem.Raw("<path d=\"M22 4L12 14.01L9 11.01\" stroke=\"#33C27F\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" />"),
		},
	}
	iconWarning := &elem.Element{
		Tag: "svg",
		Attrs: attrs.Props{
			attrs.Width:  "30",
			attrs.Height: "30",
			"viewBox":    "0 0 24 24",
			"fill":       "none",
			"xmlns":      "http://www.w3.org/2000/svg",
			attrs.Style: styles.Props{
				styles.MarginBottom: "1rem",
			}.ToInline(),
		},
		Children: []elem.Node{
			elem.Raw("<path d=\"M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z\" stroke=\"#E5993E\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" />"),
			elem.Raw("<path d=\"M12 8V12\" stroke=\"#E5993E\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" />"),
			elem.Raw("<path d=\"M12 16H12.01\" stroke=\"#E5993E\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" />"),
		},
	}
	spanIconLink := elem.Span(
		attrs.Props{
			attrs.Class: oidcIconClass,
		},
		elem.Raw("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\"><path d=\"M13.307 1H11.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L13.307 1zM12 14V8a.5.5 0 1 1 1 0v6.5a.5.5 0 0 1-.5.5H.563a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 .5-.5H8a.5.5 0 0 1 0 1H1v12h11zM4 6a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1H4zm0 2.5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1H4zM4 11a.5.5 0 1 1 0-1h5a.5.5 0 1 1 0 1H4z\" /></svg>"),
	)

	pictureUrlStyle := ""
	if strings.TrimSpace(cfg.User.GetProfilePicURL()) != "" {
		pictureUrlStyle = styles.Props{
			styles.BackgroundImage: styles.URL(cfg.User.GetProfilePicURL()),
		}.ToInline()
	}

	headerBlock := elem.Div(
		attrs.Props{
			attrs.Class: oidcHeaderBlockClass,
		},
		elem.Div(
			attrs.Props{
				attrs.Style: styles.Props{
					styles.FlexShrink: "0",
				}.ToInline(),
			},
			logo,
		),
		elem.Div(
			attrs.Props{
				attrs.Class: oidcHeaderUserBlockClass,
			},
			elem.Div(
				attrs.Props{
					attrs.Class: oidcHeaderUserBlockTruncateClass,
				},
				elem.Text(cfg.User.DisplayNameOrUsername()),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: oidcHeaderUserBlockPictureClass,
					attrs.Style: pictureUrlStyle,
				},
			),
		),
	)

	icon := iconSuccess
	if cfg.Node.IsApproved() == false {
		icon = iconWarning
	}

	titleText := fmt.Sprintf("%s successful", cfg.Verb)
	if cfg.Node.IsApproved() == false {
		titleText = fmt.Sprintf("%s, but not connected", cfg.Verb)
	}

	descriptions := elem.Section(nil)

	description := elem.P(
		attrs.Props{
			attrs.Style: styles.Props{
				styles.MarginTop: "1rem",
			}.ToInline(),
		},
		elem.Text("Your node "),
		elem.Strong(nil, elem.Text(cfg.Node.Hostname)),
		elem.Text(" is added to network, you can now close this window."),
	)
	descriptions.Children = append(descriptions.Children, description)

	if cfg.Node.IsApproved() == false {
		descriptionApproval := elem.P(
			attrs.Props{
				attrs.Style: styles.Props{
					styles.MarginTop: "1rem",
				}.ToInline(),
			},
			elem.Text("However, it can't connect until approved by the administrator a network. Once approved, your node will connect automatically."),
		)

		descriptions.Children = append(descriptions.Children, descriptionApproval)
	}

	containerTag := elem.Div(
		attrs.Props{
			attrs.Class: oidcContainerClass,
		},
		icon,
		elem.H1(
			attrs.Props{
				attrs.Class: oidcH1Class,
			},
			elem.Text(titleText),
		),
		descriptions,
		elem.Hr(attrs.Props{attrs.Class: oidcHRClass}),
		elem.H1(
			attrs.Props{
				attrs.Class: oidcH1Class,
			},
			elem.Text("Not sure how to get started?"),
		),
		elem.P(
			attrs.Props{
				attrs.Style: styles.Props{
					styles.MarginTop: "1rem",
				}.ToInline(),
			},
			elem.Text("Check out beginner and advanced guides on, or read more in the documentation."),
		),
		elem.A(
			attrs.Props{
				attrs.Class:  oidcAClass,
				attrs.Href:   "https://github.com/juanfont/headscale/tree/main/docs",
				attrs.Rel:    "noreferrer noopener",
				attrs.Target: "_blank",
			},
			spanIconLink,
			elem.Text("View the headscale documentation"),
		),
		elem.A(
			attrs.Props{
				attrs.Class:  oidcAClass,
				attrs.Href:   "https://tailscale.com/kb/",
				attrs.Rel:    "noreferrer noopener",
				attrs.Target: "_blank",
			},
			spanIconLink,
			elem.Text("View the tailscale documentation"),
		),
	)

	mainTag := elem.Main(
		attrs.Props{
			attrs.Class: oidcMainTagClass,
		},
		elem.Div(
			nil,
			headerBlock,
			containerTag,
		),
	)

	return HtmlStructure(
		elem.Head(
			nil,
			elem.Title(nil, elem.Text("Authentication Succeeded - Headscale")),
			elem.Style(attrs.Props{}),
		),
		elem.Body(
			attrs.Props{
				attrs.Class: oidcBodyTagClass,
			},
			mainTag,
		),
	).RenderWithOptions(elem.RenderOptions{StyleManager: oidcStyleMgr})
}
