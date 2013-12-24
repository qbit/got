Architecture
============
1. Dynamically loadable plugins via DSL.
2. Split components:
    - **Server component** - handles plugin processing, database connections
      defines the DSL, publishes a series of interfaces for configuring the
      bot.
    - **Client component** - actually makes connection to IRC, can run when
      server component is down.
3. General NLP package to be written that is exposed to the DSL. This will allow
   general parsing / classification of messages.
