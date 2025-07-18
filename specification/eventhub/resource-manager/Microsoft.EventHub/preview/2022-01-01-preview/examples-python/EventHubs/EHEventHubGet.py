from azure.identity import DefaultAzureCredential

from azure.mgmt.eventhub import EventHubManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventhub
# USAGE
    python eh_event_hub_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventHubManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="e2f361f0-3b27-4503-a9cc-21cfba380093",
    )

    response = client.event_hubs.get(
        resource_group_name="Default-NotificationHubs-AustraliaEast",
        namespace_name="sdk-Namespace-716",
        event_hub_name="sdk-EventHub-10",
    )
    print(response)


# x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/EventHubs/EHEventHubGet.json
if __name__ == "__main__":
    main()
