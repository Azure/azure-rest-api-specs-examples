from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservices import RecoveryServicesClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservices
# USAGE
    python put_vault_with_redundancy_settings.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesClient(
        credential=DefaultAzureCredential(),
        subscription_id="77777777-b0c6-47a2-b37c-d8e65a629c18",
    )

    response = client.vaults.begin_create_or_update(
        resource_group_name="Default-RecoveryServices-ResourceGroup",
        vault_name="swaggerExample",
        vault={
            "identity": {"type": "SystemAssigned"},
            "location": "West US",
            "properties": {
                "publicNetworkAccess": "Enabled",
                "redundancySettings": {
                    "crossRegionRestore": "Enabled",
                    "standardTierStorageRedundancy": "GeoRedundant",
                },
            },
            "sku": {"name": "Standard"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PUTVault_WithRedundancySettings.json
if __name__ == "__main__":
    main()
