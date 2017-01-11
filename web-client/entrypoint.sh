# We run again yarn install on the entrypoint in case the user modified package.json or yarn.lock, we want them to see the effect
# right away.
# The node_modules should already be pre-loaded in the built image, so this command should be fast.
#yarn install --pure-lockfile
exec "$@"