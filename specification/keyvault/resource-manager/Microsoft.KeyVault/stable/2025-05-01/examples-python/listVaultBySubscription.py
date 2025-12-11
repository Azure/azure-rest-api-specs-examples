from azure.identity import DefaultAzureCredential

from azure.mgmt.keyvault import KeyVaultManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-keyvault
# USAGE
    python list_vault_by_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KeyVaultManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.vaults.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: 2025-05-01/listVaultBySubscription.json
if __name__ == "__main__":
    main()
