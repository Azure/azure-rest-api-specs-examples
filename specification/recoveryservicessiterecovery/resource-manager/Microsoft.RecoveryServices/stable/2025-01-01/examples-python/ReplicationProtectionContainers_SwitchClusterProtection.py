from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_protection_containers_switch_cluster_protection.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="7c943c1b-5122-4097-90c8-861411bdd574",
        resource_group_name="resourceGroupPS1",
        resource_name="vault1",
    )

    response = client.replication_protection_containers.begin_switch_cluster_protection(
        resource_name="vault1",
        fabric_name="fabric-pri-eastus",
        protection_container_name="pri-cloud-eastus",
        switch_input={
            "properties": {
                "providerSpecificDetails": {
                    "instanceType": "A2A",
                    "policyId": "/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/klncksan",
                    "protectedItemsDetail": [
                        {
                            "recoveryResourceGroupId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr",
                            "replicationProtectedItemName": "yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio",
                            "vmManagedDisks": [
                                {
                                    "diskId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql0-osdisk",
                                    "primaryStagingAzureStorageAccountId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache",
                                    "recoveryResourceGroupId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr",
                                }
                            ],
                        },
                        {
                            "recoveryResourceGroupId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr",
                            "replicationProtectedItemName": "kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4",
                            "vmManagedDisks": [
                                {
                                    "diskId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql1-osdisk",
                                    "primaryStagingAzureStorageAccountId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache",
                                    "recoveryResourceGroupId": "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr",
                                }
                            ],
                        },
                    ],
                    "recoveryContainerId": "/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-rec-westus/replicationProtectionContainers/rec-cloud-westus",
                },
                "replicationProtectionClusterName": "testcluster",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionContainers_SwitchClusterProtection.json
if __name__ == "__main__":
    main()
