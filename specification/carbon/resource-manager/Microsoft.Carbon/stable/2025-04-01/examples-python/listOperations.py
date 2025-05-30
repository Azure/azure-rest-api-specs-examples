from azure.identity import DefaultAzureCredential

from azure.mgmt.carbonoptimization import CarbonOptimizationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-carbonoptimization
# USAGE
    python list_operations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CarbonOptimizationMgmtClient(
        credential=DefaultAzureCredential(),
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: 2025-04-01/listOperations.json
if __name__ == "__main__":
    main()
