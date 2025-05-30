from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_restore_database_account_create_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.database_accounts.begin_create_or_update(
        resource_group_name="rg1",
        account_name="ddb1",
        create_update_parameters={
            "kind": "GlobalDocumentDB",
            "location": "westus",
            "properties": {
                "apiProperties": {"serverVersion": "3.2"},
                "backupPolicy": {"continuousModeProperties": {"tier": "Continuous30Days"}, "type": "Continuous"},
                "consistencyPolicy": {
                    "defaultConsistencyLevel": "BoundedStaleness",
                    "maxIntervalInSeconds": 10,
                    "maxStalenessPrefix": 200,
                },
                "createMode": "Restore",
                "databaseAccountOfferType": "Standard",
                "enableAnalyticalStorage": True,
                "enableFreeTier": False,
                "enableMaterializedViews": False,
                "keyVaultKeyUri": "https://myKeyVault.vault.azure.net",
                "locations": [{"failoverPriority": 0, "isZoneRedundant": False, "locationName": "southcentralus"}],
                "minimalTlsVersion": "Tls",
                "restoreParameters": {
                    "databasesToRestore": [
                        {"collectionNames": ["collection1", "collection2"], "databaseName": "db1"},
                        {"collectionNames": ["collection3", "collection4"], "databaseName": "db2"},
                    ],
                    "restoreMode": "PointInTime",
                    "restoreSource": "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc",
                    "restoreTimestampInUtc": "2021-03-11T22:05:09Z",
                    "restoreWithTtlDisabled": False,
                    "sourceBackupLocation": "westus",
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBRestoreDatabaseAccountCreateUpdate.json
if __name__ == "__main__":
    main()
