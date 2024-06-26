from azure.identity import DefaultAzureCredential
from azure.mgmt.selfhelp import SelfHelpMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-selfhelp
# USAGE
    python troubleshooter_restart.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SelfHelpMgmtClient(
        credential=DefaultAzureCredential(),
    )

    response = client.troubleshooters.restart(
        scope="subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
        troubleshooter_name="abf168ed-1b54-454a-86f6-e4b62253d3b1",
    )
    print(response)


# x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Troubleshooter_Restart.json
if __name__ == "__main__":
    main()
