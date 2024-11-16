from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridcompute import HybridComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcompute
# USAGE
    python gateway_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffd506c8-3415-42d3-9612-fdb423fb17df",
    )

    response = client.gateways.begin_create_or_update(
        resource_group_name="myResourceGroup",
        gateway_name="{gatewayName}",
        parameters={"location": "eastus2euap", "properties": {"allowedFeatures": ["*"], "gatewayType": "Public"}},
    ).result()
    print(response)


# x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_CreateOrUpdate.json
if __name__ == "__main__":
    main()
