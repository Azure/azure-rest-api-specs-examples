from azure.identity import DefaultAzureCredential
from azure.mgmt.quantum import AzureQuantumManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-quantum
# USAGE
    python operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureQuantumManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/operations.json
if __name__ == "__main__":
    main()
