package day12

import "testing"

func TestPossiblePaths(t *testing.T) {
	actual := PossiblePaths("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end")
	expected := 10
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestPossiblePaths2(t *testing.T) {
	actual := PossiblePaths("dc-end\nHN-start\nstart-kj\ndc-start\ndc-HN\nLN-dc\nHN-end\nkj-sa\nkj-HN\nkj-dcdc-end\nHN-start\nstart-kj\ndc-start\ndc-HN\nLN-dc\nHN-end\nkj-sa\nkj-HN\nkj-dc")
	expected := 19
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestPossiblePaths3(t *testing.T) {
	actual := PossiblePaths("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW")
	expected := 226
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestPossiblePaths_Real(t *testing.T) {
	actual := PossiblePaths("HF-qu\nend-CF\nCF-ae\nvi-HF\nvt-HF\nqu-CF\nhu-vt\nCF-pk\nCF-vi\nqu-ae\nae-hu\nHF-start\nvt-end\nae-HF\nend-vi\nvi-vt\nhu-start\nstart-ae\nCS-hu\nCF-vt")
	expected := 3802
	if actual != expected {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
