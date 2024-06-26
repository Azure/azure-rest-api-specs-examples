from azure.identity import DefaultAzureCredential
from azure.mgmt.quantum import AzureQuantumManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-quantum
# USAGE
    python quantum_workspaces_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureQuantumManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.workspaces.begin_create_or_update(
        resource_group_name="quantumResourcegroup",
        workspace_name="quantumworkspace1",
        quantum_workspace={
            "location": "West US",
            "properties": {
                "providers": [
                    {"providerId": "Honeywell", "providerSku": "Basic"},
                    {"providerId": "IonQ", "providerSku": "Basic"},
                    {"providerId": "OneQBit", "providerSku": "Basic"},
                ],
                "storageAccount": "/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesPut.json
if __name__ == "__main__":
    main()
