from azure.identity import DefaultAzureCredential

from azure.mgmt.eventgrid import EventGridManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventgrid
# USAGE
    python topics_list_event_types.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventGridManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="8f6b6269-84f2-4d09-9e31-1127efcd1e40",
    )

    response = client.topics.list_event_types(
        resource_group_name="examplerg",
        provider_namespace="Microsoft.Storage",
        resource_type_name="storageAccounts",
        resource_name="ExampleStorageAccount",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Topics_ListEventTypes.json
if __name__ == "__main__":
    main()
