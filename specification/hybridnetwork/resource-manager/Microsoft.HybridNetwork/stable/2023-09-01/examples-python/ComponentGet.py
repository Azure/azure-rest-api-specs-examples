from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python component_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.components.get(
        resource_group_name="rg",
        network_function_name="testNf",
        component_name="testComponent",
    )
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ComponentGet.json
if __name__ == "__main__":
    main()
