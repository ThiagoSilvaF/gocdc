module github.com/ThiagoSilvaF/gocdc

go 1.13

require github.com/lib/pq v1.3.0

replace (
	gocdc/kafka v0.0.0 => ./gocdc/kafka
	gocdc/postgres v0.0.0 => ./gocdc/postgres
)
