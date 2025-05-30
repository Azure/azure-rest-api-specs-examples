from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_database_account_create_max.py

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
            "identity": {
                "type": "SystemAssigned,UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {}
                },
            },
            "kind": "MongoDB",
            "location": "westus",
            "properties": {
                "analyticalStorageConfiguration": {"schemaType": "WellDefined"},
                "apiProperties": {"serverVersion": "3.2"},
                "backupPolicy": {
                    "periodicModeProperties": {
                        "backupIntervalInMinutes": 240,
                        "backupRetentionIntervalInHours": 8,
                        "backupStorageRedundancy": "Geo",
                    },
                    "type": "Periodic",
                },
                "capacity": {"totalThroughputLimit": 2000},
                "consistencyPolicy": {
                    "defaultConsistencyLevel": "BoundedStaleness",
                    "maxIntervalInSeconds": 10,
                    "maxStalenessPrefix": 200,
                },
                "cors": [{"allowedOrigins": "https://test"}],
                "createMode": "Default",
                "databaseAccountOfferType": "Standard",
                "defaultIdentity": "FirstPartyIdentity",
                "enableAnalyticalStorage": True,
                "enableBurstCapacity": True,
                "enableFreeTier": False,
                "enablePerRegionPerPartitionAutoscale": True,
                "ipRules": [{"ipAddressOrRange": "23.43.230.120"}, {"ipAddressOrRange": "110.12.240.0/12"}],
                "isVirtualNetworkFilterEnabled": True,
                "keyVaultKeyUri": "https://myKeyVault.vault.azure.net",
                "locations": [
                    {"failoverPriority": 0, "isZoneRedundant": False, "locationName": "southcentralus"},
                    {"failoverPriority": 1, "isZoneRedundant": False, "locationName": "eastus"},
                ],
                "minimalTlsVersion": "Tls12",
                "networkAclBypass": "AzureServices",
                "networkAclBypassResourceIds": [
                    "/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName"
                ],
                "publicNetworkAccess": "Enabled",
                "virtualNetworkRules": [
                    {
                        "id": "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
                        "ignoreMissingVNetServiceEndpoint": False,
                    }
                ],
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBDatabaseAccountCreateMax.json
if __name__ == "__main__":
    main()
