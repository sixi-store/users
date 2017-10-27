{{define "user.list"}}
	select
		*
	from
		users
	where 1
		{{if .name}}
		    and name = :name
		{{end}}
		{{if .sex}}
            and sex = :sex
        {{end}}
        {{if .limit}}
            limit :offset, :limit
        {{end}}
{{end}}

{{define "user.count"}}
	select
		count(*) as count
	from
		users
	where 1
		{{if .name}}
		    and name = :name
		{{end}}
		{{if .sex}}
            and sex = :sex
        {{end}}
{{end}}


{{define "user.info"}}
	select
		*
	from
		users
	where 1
		{{if .name}}
		    and name = :name
		{{end}}
	limit 1
{{end}}