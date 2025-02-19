from azure.identity import DefaultAzureCredential

from azure.mgmt.loganalytics import LogAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loganalytics
# USAGE
    python workspaces_saved_searches_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-00000000000",
    )

    response = client.saved_searches.create_or_update(
        resource_group_name="TestRG",
        workspace_name="TestWS",
        saved_search_id="00000000-0000-0000-0000-00000000000",
        parameters={
            "properties": {
                "category": "Saved Search Test Category",
                "displayName": "Create or Update Saved Search Test",
                "functionAlias": "heartbeat_func",
                "functionParameters": "a:int=1",
                "query": "Heartbeat | summarize Count() by Computer | take a",
                "tags": [{"name": "Group", "value": "Computer"}],
                "version": 2,
            }
        },
    )
    print(response)


# x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesSavedSearchesCreateOrUpdate.json
if __name__ == "__main__":
    main()
