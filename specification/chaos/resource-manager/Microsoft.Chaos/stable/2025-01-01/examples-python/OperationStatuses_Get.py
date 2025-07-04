from azure.identity import DefaultAzureCredential

from azure.mgmt.chaos import ChaosManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-chaos
# USAGE
    python operation_statuses_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ChaosManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="e25c0d12-0335-4fec-8ef8-3b4f9a10649e",
    )

    response = client.operation_statuses.get(
        location="westus2",
        operation_id="4bdadd97-207c-4de8-9bba-08339ae099c7",
    )
    print(response)


# x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2025-01-01/examples/OperationStatuses_Get.json
if __name__ == "__main__":
    main()
