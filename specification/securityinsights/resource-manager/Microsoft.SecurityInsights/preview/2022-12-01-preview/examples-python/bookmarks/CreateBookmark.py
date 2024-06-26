from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_bookmark.py

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

    response = client.bookmarks.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        bookmark_id="73e01a99-5cd7-4139-a149-9f2736ff2ab5",
        bookmark={
            "etag": '"0300bf09-0000-0000-0000-5c37296e0000"',
            "properties": {
                "created": "2021-09-01T13:15:30Z",
                "createdBy": {"objectId": "2046feea-040d-4a46-9e2b-91c2941bfa70"},
                "displayName": "My bookmark",
                "entityMappings": [
                    {
                        "entityType": "Account",
                        "fieldMappings": [{"identifier": "Fullname", "value": "johndoe@microsoft.com"}],
                    }
                ],
                "labels": ["Tag1", "Tag2"],
                "notes": "Found a suspicious activity",
                "query": "SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)",
                "queryResult": "Security Event query result",
                "tactics": ["Execution"],
                "techniques": ["T1609"],
                "updated": "2021-09-01T13:15:30Z",
                "updatedBy": {"objectId": "2046feea-040d-4a46-9e2b-91c2941bfa70"},
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/bookmarks/CreateBookmark.json
if __name__ == "__main__":
    main()
