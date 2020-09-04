// This file is automatically generated. Do not modify it manually.

const manifest = JSON.parse(`
{
    "id": "com.mattermost.webapp-chr-creation",
    "name": "CHR Creation Plugin",
    "description": "This Plugin can be used to create Concept Help Request from Mattermost",
    "version": "0.1.0",
    "min_server_version": "5.12.0",
    "server": {
        "executables": {
            "linux-amd64": "server/dist/plugin-linux-amd64",
            "darwin-amd64": "server/dist/plugin-darwin-amd64",
            "windows-amd64": "server/dist/plugin-windows-amd64.exe"
        },
        "executable": ""
    },
    "webapp": {
        "bundle_path": "webapp/dist/main.js"
    },
    "settings_schema": {
        "header": "The settings page of CHR creation plugin",
        "footer": "Made with \u003c3 by Scaler",
        "settings": [
            {
                "key": "ChrTriggerWords",
                "display_name": "CHR Trigger words",
                "type": "longtext",
                "help_text": "The words to be considered as Triggers for the CHR creation bot",
                "placeholder": "",
                "default": "Doubt ? what when how would"
            }
        ]
    }
}
`);

export default manifest;
export const id = manifest.id;
export const version = manifest.version;
