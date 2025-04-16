package lint


var SeverityMap = map[string]struct {
    Emoji  string
    Weight int
}{
    "low": {
        Emoji:  "🟡",
        Weight: 2,
    },
    "medium": {
        Emoji:  "🟠",
        Weight: 5,
    },
    "high": {
        Emoji:  "🔴",
        Weight: 8,
    },
    "critical": {
        Emoji:  "🚨",
        Weight: 10,
    },
}
