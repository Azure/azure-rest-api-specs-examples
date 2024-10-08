from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsightcontainers import HDInsightContainersMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsightcontainers
# USAGE
    python cluster_upgrade_rollback.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HDInsightContainersMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="10e32bab-26da-4cc4-a441-52b318f824e6",
    )

    response = client.clusters.begin_upgrade_manual_rollback(
        resource_group_name="hiloResourcegroup",
        cluster_pool_name="clusterpool1",
        cluster_name="cluster1",
        cluster_rollback_upgrade_request={
            "properties": {
                "upgradeHistory": "/subscriptions/10e32bab-26da-4cc4-a441-52b318f824e6/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterpools/clusterpool1/clusters/cluster1/upgradeHistories/01_11_2024_02_35_03_AM-HotfixUpgrade"
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ClusterUpgradeRollback.json
if __name__ == "__main__":
    main()
