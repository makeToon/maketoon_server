package config

// errorFormat : server logger format
func ErrorFormat() string {
	return `{
		time: ${time_rfc3339_nano}
		id: ${id}
		remote_ip: ${remote_ip}
		host: ${host}
		method: ${method}
		uri: ${uri}
		status: ${status}
		error: ${error}
		latency: ${latency}
		latency_human: ${latency_human}
		bytes_in: ${bytes_in}
		bytes_out: ${bytes_out}
	}`
}