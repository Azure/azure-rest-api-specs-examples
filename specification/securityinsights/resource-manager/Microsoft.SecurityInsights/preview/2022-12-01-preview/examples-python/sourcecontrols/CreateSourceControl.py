from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_source_control.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.source_controls.create(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        source_control_id="789e0c1f-4a3d-43ad-809c-e713b677b04a",
        source_control={
            "etag": '"0300bf09-0000-0000-0000-5c37296e0000"',
            "properties": {
                "contentTypes": ["AnalyticRules", "Workbook"],
                "description": "This is a source control",
                "displayName": "My Source Control",
                "repoType": "Github",
                "repository": {
                    "branch": "master",
                    "displayUrl": "https://github.com/user/repo",
                    "pathMapping": [
                        {"contentType": "AnalyticRules", "path": "path/to/rules"},
                        {"contentType": "Workbook", "path": "path/to/workbooks"},
                    ],
                    "url": "https://github.com/user/repo",
                },
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/sourcecontrols/CreateSourceControl.json
if __name__ == "__main__":
    main()
