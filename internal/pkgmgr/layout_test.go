package pkgmgr

import (
	"path/filepath"
	"runtime"
	"testing"
)

// TestHomeHonorsPIDANHOME verifies Home prefers PIDAN_HOME over the default.
func TestHomeHonorsPIDANHOME(t *testing.T) {
	t.Setenv("PIDAN_HOME", "/custom/pidan")
	if got := Home(); got != "/custom/pidan" {
		t.Errorf("Home() = %q, want /custom/pidan", got)
	}
}

// TestTypeDirsUnderHome verifies the plugins/commands/themes dirs nest under
// $PIDAN_HOME.
func TestTypeDirsUnderHome(t *testing.T) {
	t.Setenv("PIDAN_HOME", "/custom/pidan")
	cases := map[PackageType]string{
		TypeExtension: filepath.Join("/custom/pidan", "plugins"),
		TypePrompt:    filepath.Join("/custom/pidan", "commands"),
		TypeTheme:     filepath.Join("/custom/pidan", "themes"),
	}
	for typ, want := range cases {
		if got := DirForType(typ); got != want {
			t.Errorf("DirForType(%s) = %q, want %q", typ, got, want)
		}
	}
}

// TestSkillsDirHonorsOverride verifies skills use PIDAN_SKILLS_DIR, not $PIDAN_HOME.
func TestSkillsDirHonorsOverride(t *testing.T) {
	t.Setenv("PIDAN_SKILLS_DIR", "/custom/skills")
	if got := SkillsDir(); got != "/custom/skills" {
		t.Errorf("SkillsDir() = %q, want /custom/skills", got)
	}
	if got := DirForType(TypeSkill); got != "/custom/skills" {
		t.Errorf("DirForType(skill) = %q, want /custom/skills", got)
	}
}

// TestSkillsDirDefault verifies skills default to ~/.agents/skills.
func TestSkillsDirDefault(t *testing.T) {
	t.Setenv("PIDAN_SKILLS_DIR", "")
	// os.UserHomeDir reads HOME on unix and USERPROFILE on windows.
	home := "/home/tester"
	if runtime.GOOS == "windows" {
		t.Setenv("USERPROFILE", home)
	} else {
		t.Setenv("HOME", home)
	}
	want := filepath.Join(home, ".agents", "skills")
	if got := SkillsDir(); got != want {
		t.Errorf("SkillsDir() = %q, want %q", got, want)
	}
}

// TestDirForUnknownType verifies an unknown type yields "".
func TestDirForUnknownType(t *testing.T) {
	if got := DirForType(PackageType("bogus")); got != "" {
		t.Errorf("DirForType(bogus) = %q, want empty", got)
	}
}
