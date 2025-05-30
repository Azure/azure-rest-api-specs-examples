from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python cluster_recovery_points_list_by_replication_protection_cluster.py

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

    response = client.cluster_recovery_points.list_by_replication_protection_cluster(
        resource_name="vault1",
        fabric_name="fabric-pri-eastus",
        protection_container_name="pri-cloud-eastus",
        replication_protection_cluster_name="testcluster",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ClusterRecoveryPoints_ListByReplicationProtectionCluster.json
if __name__ == "__main__":
    main()
