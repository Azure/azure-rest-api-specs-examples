from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.dataprotection import DataProtectionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dataprotection
# USAGE
    python put_backup_vault_with_cmk.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataProtectionMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="0b352192-dcac-4cc7-992e-a96190ccc68c",
    )

    response = client.backup_vaults.begin_create_or_update(
        resource_group_name="SampleResourceGroup",
        vault_name="swaggerExample",
        parameters={
            "identity": {"type": "None"},
            "location": "WestUS",
            "properties": {
                "monitoringSettings": {"azureMonitorAlertSettings": {"alertsForAllJobFailures": "Enabled"}},
                "securitySettings": {
                    "encryptionSettings": {
                        "infrastructureEncryption": "Enabled",
                        "kekIdentity": {
                            "identityId": "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi",
                            "identityType": "UserAssigned",
                        },
                        "keyVaultProperties": {
                            "keyUri": "https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3"
                        },
                        "state": "Enabled",
                    },
                    "immutabilitySettings": {"state": "Disabled"},
                    "softDeleteSettings": {"retentionDurationInDays": 0, "state": "Off"},
                },
                "storageSettings": [{"datastoreType": "VaultStore", "type": "LocallyRedundant"}],
            },
            "tags": {"key1": "val1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/PutBackupVaultWithCMK.json
if __name__ == "__main__":
    main()
