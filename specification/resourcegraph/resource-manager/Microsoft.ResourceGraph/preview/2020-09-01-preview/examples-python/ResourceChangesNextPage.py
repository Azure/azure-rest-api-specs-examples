from azure.identity import DefaultAzureCredential
from azure.mgmt.resourcegraph import ResourceGraphClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcegraph
# USAGE
    python resource_changes_next_page.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceGraphClient(
        credential=DefaultAzureCredential(),
    )

    response = client.resource_changes(
        parameters={
            "$skipToken": "ew0KICAiJGlkIjogIjEiLA0KICAiRW5kVGltZSI6ICJcL0RhdGUoMTU1MDc0NT",
            "$top": 2,
            "interval": {"end": "2018-10-31T12:09:03.141Z", "start": "2018-10-30T12:09:03.141Z"},
            "resourceIds": [
                "/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"
            ],
        },
    )
    print(response)


# x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2020-09-01-preview/examples/ResourceChangesNextPage.json
if __name__ == "__main__":
    main()
