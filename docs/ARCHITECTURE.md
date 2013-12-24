Architecture
============
1. Dynamically loadable plugins via DSL.
2. Split components:
	a. Server component - handles plugin processing, database connections
	defines the DSL, publishes a series of interfaces for configuring the
	bot.
	b. Client component - actually makes connection to IRC, can run when
	server component is down.
