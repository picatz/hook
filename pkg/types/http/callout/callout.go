package callout

type Callback = func(numHeaders, bodySize, numTrailers int)
