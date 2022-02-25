package internal

import (
	"context"
	"fmt"
	"io"

	"github.com/ajstarks/svgo"
)

const (
	width        = 1200
	height       = 630
	paddingX     = 100
	paddingY     = 150
	fontFamily   = "sans-serif"
	bgColorStart = "#171717"
	bgColorEnd   = "#6d28d9"

	titleSize  = 50
	titleColor = "#f5f5f5"

	profileImageSize   = 60
	profileImageMargin = 20

	authorNameSize  = 30
	authorNameColor = "#f5f5f5"
)

func GenerateSVG(ctx context.Context, payload *Payload) io.Reader {
	pr, pw := io.Pipe()
	go func(_ctx context.Context) {
		defer pw.Close()

		canvas := svg.New(pw)
		canvas.Start(width, height)

		// Background
		gradientID := "gradient"
		canvas.Def()
		canvas.LinearGradient(gradientID, 0, 70, 100, 30, []svg.Offcolor{
			{Offset: 5, Color: bgColorStart, Opacity: 100},
			{Offset: 95, Color: bgColorEnd, Opacity: 100},
		})
		canvas.DefEnd()
		canvas.Rect(0, 0, width, height, fmt.Sprintf("fill: url(#%s)", gradientID))

		// Title
		titleFirst, titleSecond, titleMore := splitTitle(payload.Title)
		titleAttr := fmt.Sprintf("font-size: %dpt; fill: %s; font-family: %s; font-weight: 700", titleSize, titleColor, fontFamily)
		canvas.Text(paddingX, paddingY+titleSize, titleFirst, titleAttr)
		if titleSecond != "" {
			if titleMore {
				titleSecond = titleSecond + "..."
			}
			canvas.Text(paddingX, paddingY+titleSize*3, titleSecond, titleAttr)
		}

		// Profile image
		profileImageID := "profile-image"
		profileImageAttr := `preserveAspectRatio="xMidYMid slice"`
		canvas.Def()
		canvas.Pattern(profileImageID, 0, 0, profileImageSize, profileImageSize, "userSpaceOnUse")
		canvas.Image(0, 0, profileImageSize, profileImageSize, payload.ProfileImageURL, profileImageAttr)
		canvas.PatternEnd()
		canvas.DefEnd()

		profileImageX := paddingX + profileImageSize/2
		profileImageY := height - (paddingY + profileImageSize/2)
		canvas.Circle(profileImageX, profileImageY, profileImageSize/2, `fill="url(#`+profileImageID+`)"`)

		// Author name
		authorNameX := paddingX + profileImageSize + profileImageMargin
		authorNameY := height - (paddingY + (profileImageSize-authorNameSize)/2)
		authorNameAttr := fmt.Sprintf("font-size: %dpt; fill: %s; font-family: %s; font-weight: 300", authorNameSize, authorNameColor, fontFamily)
		canvas.Text(authorNameX, authorNameY, payload.AuthorName, authorNameAttr)

		canvas.End()
	}(ctx)
	return pr
}

func splitTitle(title string) (first string, second string, more bool) {
	lineSize := 26
	first, left := splitText(title, lineSize)
	second, left = splitText(left, lineSize)
	more = left != ""
	return
}
