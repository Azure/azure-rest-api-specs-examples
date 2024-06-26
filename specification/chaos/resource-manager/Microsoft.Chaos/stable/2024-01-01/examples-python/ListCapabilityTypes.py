from azure.identity import DefaultAzureCredential
from azure.mgmt.chaos import ChaosManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-chaos
# USAGE
    python list_capability_types.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ChaosManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="6b052e15-03d3-4f17-b2e1-be7f07588291",
    )

    response = client.capability_types.list(
        location_name="westus2",
        target_type_name="Microsoft-VirtualMachine",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/ListCapabilityTypes.json
if __name__ == "__main__":
    main()
