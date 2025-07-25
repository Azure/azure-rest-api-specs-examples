from azure.identity import DefaultAzureCredential

from azure.mgmt.eventhub import EventHubManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventhub
# USAGE
    python private_end_point_connection_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventHubManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subID",
    )

    response = client.private_endpoint_connections.list(
        resource_group_name="SDK-EventHub-4794",
        namespace_name="sdk-Namespace-5828",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/NameSpaces/PrivateEndPointConnectionList.json
if __name__ == "__main__":
    main()
