from azure.identity import DefaultAzureCredential

from azure.mgmt.digitaltwins import AzureDigitalTwinsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-digitaltwins
# USAGE
    python private_endpoint_connection_put_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureDigitalTwinsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="50016170-c839-41ba-a724-51e9df440b9e",
    )

    response = client.private_endpoint_connections.begin_create_or_update(
        resource_group_name="resRg",
        resource_name="myDigitalTwinsService",
        private_endpoint_connection_name="myPrivateConnection",
        private_endpoint_connection={
            "properties": {
                "privateLinkServiceConnectionState": {
                    "description": "Approved by johndoe@company.com.",
                    "status": "Approved",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateEndpointConnectionPut_example.json
if __name__ == "__main__":
    main()
