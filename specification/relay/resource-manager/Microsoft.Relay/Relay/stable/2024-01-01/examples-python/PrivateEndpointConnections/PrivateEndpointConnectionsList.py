from azure.identity import DefaultAzureCredential

from azure.mgmt.relay import RelayAPIMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-relay
# USAGE
    python private_endpoint_connections_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RelayAPIMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.private_endpoint_connections.list(
        resource_group_name="myResourceGroup",
        namespace_name="example-RelayNamespace-5849",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2024-01-01/PrivateEndpointConnections/PrivateEndpointConnectionsList.json
if __name__ == "__main__":
    main()
