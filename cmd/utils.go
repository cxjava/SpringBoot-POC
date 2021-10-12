package cmd

import "net/url"

func concatUrl(base, ref string) (string, error) {
	b, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	obj, err := b.Parse(ref)
	if err != nil {
		return "", err
	}
	return obj.String(), nil
}
