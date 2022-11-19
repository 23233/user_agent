package user_agent

import "strings"

func listSectionGetName(sections []section, name string) (section, bool) {
	for _, sec := range sections {
		if sec.name == name {
			return sec, true
		}
	}
	return section{}, false
}

func arraySectionFilter(sections []section, filter string, fuzzy bool) (section, int, bool, *string) {
	for index, sec := range sections {
		if fuzzy {
			var ab strings.Builder
			ab.WriteString(sec.name)
			ab.WriteString(sec.version)

			if strings.Contains(ab.String(), filter) {
				return sec, index, true, nil
			}

			for _, com := range sec.comment {
				if strings.Contains(com, filter) {
					return sec, index, true, &com
				}
			}

		} else {
			if sec.name == filter || sec.version == filter {
				return sec, index, true, nil
			}
			for _, com := range sec.comment {
				if com == filter {
					return sec, index, true, &com
				}
			}
		}
	}
	return section{}, 0, false, nil
}
