SubzellePluginInfo = provider(fields = {
    "name": "proto plugin name",
    "address": "plugin gRPC address",
    "executable": "plugin tool executable",
    "data": "additional data",
})

def _subzelle_plugin_impl(ctx):
    return [SubzellePluginInfo(
        data = ctx.files.data,
        executable = ctx.executable.executable,
        name = ctx.label.name,
        address = ctx.attr.address,
    )]

subzelle_plugin = rule(
    implementation = _subzelle_plugin_impl,
    attrs = {
        "executable": attr.label(
            doc = "The plugin binary",
            cfg = "host",
            allow_files = True,
            executable = True,
            mandatory = True,
        ),
        "address": attr.string(
            doc = "Optional network address for subzelle controller plugin to connect to",
        ),
        "data": attr.label_list(
            doc = "Additional files that should travel with the plugin",
            allow_files = True,
        ),
    },
)
