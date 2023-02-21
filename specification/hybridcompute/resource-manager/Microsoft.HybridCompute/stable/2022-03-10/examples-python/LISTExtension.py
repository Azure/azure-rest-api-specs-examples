from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridcompute import HybridComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcompute
# USAGE
    python list_extension.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscriptionId}",
    )

    response = client.machine_extensions.list(
        resource_group_name="myResourceGroup",
        machine_name="myMachine",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-03-10/examples/LISTExtension.json
if __name__ == "__main__":
    main()
