from azure.identity import DefaultAzureCredential
from azure.mgmt.relay import RelayAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-relay
# USAGE
    python relay_hybridconnection_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RelayAPI(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.hybrid_connections.delete(
        resource_group_name="resourcegroup",
        namespace_name="example-RelayNamespace-01",
        hybrid_connection_name="example-Relay-Hybrid-01",
    )
    print(response)


# x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridconnectionDelete.json
if __name__ == "__main__":
    main()
