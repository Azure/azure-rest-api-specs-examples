from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_migration_items_test_migrate.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="cb53d0c3-bd59-4721-89bc-06916a9147ef",
        resource_group_name="resourcegroup1",
        resource_name="migrationvault",
    )

    response = client.replication_migration_items.begin_test_migrate(
        fabric_name="vmwarefabric1",
        protection_container_name="vmwareContainer1",
        migration_item_name="virtualmachine1",
        test_migrate_input={
            "properties": {
                "providerSpecificDetails": {
                    "instanceType": "VMwareCbt",
                    "networkId": "/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1",
                    "osUpgradeVersion": "Microsoft Windows Server 2019 (64-bit)",
                    "recoveryPointId": "/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationFabrics/vmwarefabric1/replicationProtectionContainers/vmwareContainer1/replicationMigrationItems/virtualmachine1/migrationRecoveryPoints/9e737191-317e-43d0-8c83-e32ac3b34686",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationMigrationItems_TestMigrate.json
if __name__ == "__main__":
    main()
