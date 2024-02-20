package goprogress

import "errors"

var ErrorProgressOverTotalCount = errors.New("cannot increment progressover the total count")
