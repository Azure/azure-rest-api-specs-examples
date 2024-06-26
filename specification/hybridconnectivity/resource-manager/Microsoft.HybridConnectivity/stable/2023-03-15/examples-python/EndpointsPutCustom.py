from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridconnectivity import HybridConnectivityMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridconnectivity
# USAGE
    python endpoints_put_custom.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridConnectivityMgmtClient(
        credential=DefaultAzureCredential(),
    )

    response = client.endpoints.create_or_update(
        resource_uri="subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine",
        endpoint_name="custom",
        endpoint_resource={
            "properties": {
                "resourceId": "/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.Relay/namespaces/custom-relay-namespace",
                "type": "custom",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsPutCustom.json
if __name__ == "__main__":
    main()
