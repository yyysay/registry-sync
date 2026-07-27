package planner

import (
	"fmt"
	"strings"

	"registry-sync/model"
)

func Dump(plans []model.Plan) string {

	var b strings.Builder

	b.WriteString("同步计划\n")
	b.WriteString("====================================================================\n\n")

	if len(plans) == 0 {
		return b.String()
	}

	b.WriteString("COPY PLAN\n")
	b.WriteString("--------------------------------------------------------------------\n")

	b.WriteString(
		fmt.Sprintf(
			"%-45s %-55s %-20s\n",
			"SOURCE",
			"TARGET",
			"PLATFORM",
		),
	)

	b.WriteString(
		"--------------------------------------------------------------------\n",
	)

	for _, plan := range plans {

		source := buildSourceName(plan.Image)

		platform := ""

		if len(plan.Image.Platform) > 0 {
			platform = strings.Join(
				plan.Image.Platform,
				",",
			)
		}

		for _, target := range plan.Targets {

			b.WriteString(
				fmt.Sprintf(
					"%-45s %-55s %-20s\n",
					source,
					buildTargetName(
						plan.Image,
						target,
					),
					platform,
				),
			)
		}
	}

	return b.String()
}

func buildSourceName(
	image model.Image,
) string {

	result := image.Registry + "/" + image.Repository

	if image.Tag != "" {
		result += ":" + image.Tag
	}

	return result
}

func buildTargetName(
	image model.Image,
	target model.Target,
) string {

	repository := image.Repository

	if target.Flatten {

		parts := strings.Split(
			repository,
			"/",
		)

		repository = parts[len(parts)-1]
	}

	if target.Namespace != "" {

		repository =
			target.Namespace +
				"/" +
				repository
	}

	result := target.Registry + "/" + repository

	if image.Tag != "" {
		result += ":" + image.Tag
	}

	return result
}
