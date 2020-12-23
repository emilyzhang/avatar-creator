package avatar

import (
	"image/color"
	"path/filepath"
	"strings"
)

// Avatar is composed of a set of features
// some features are required (eyes), others are optional (accessories, glasses, etc.)
// each feature may have 1 or multiple layers
// each layer has a set ordering in the image (for now)
// some features have different options that you can choose (only one of them)

type Avatar struct {
	features []Feature
}

type Feature struct {
	name         string
	optionsExist bool
	opacity      int // number from 1 to 100
}

type Layer struct {
	feature       Feature
	featureOption string
	overlay       color.Color
	opacity       int
}

func (l Layer) path() string {
	components := strings.Join([]string{
		string(l.feature.name),
		l.featureOption,
	}, "_")
	return filepath.Join("layers", components)
}

var (
	faceOutline = Feature{
		name:    "face_outline",
		opacity: 100,
	}
	eyeColor = Feature{
		name:    "eye_color",
		opacity: 100,
	}
	eyeShadow = Feature{
		name:    "eye_shadow",
		opacity: 100,
	}
	eyePupil = Feature{
		name:    "eye_pupil",
		opacity: 100,
	}
	eyeOutline = Feature{
		name:    "eye_outline",
		opacity: 100,
	}
	eyeLashes = Feature{
		name:    "eye_lashes",
		opacity: 100,
	}
	hairColor = Feature{
		name:         "hair_color",
		opacity:      100,
		optionsExist: true,
	}
	hairOutline = Feature{
		name:         "hair_outline",
		opacity:      100,
		optionsExist: true,
	}
	lipsOutline = Feature{
		name:    "lips_outline",
		opacity: 100,
	}

	featureOrder = []Feature{
		faceOutline,
		eyeColor,
		eyeShadow,
		eyePupil,
		eyeOutline,
		eyeLashes,
		hairColor,
		hairOutline,
		lipsOutline,
	}
)
