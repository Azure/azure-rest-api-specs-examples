from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsightcontainers import HDInsightContainersMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsightcontainers
# USAGE
    python upgrade_hotfix_for_cluster.py

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

    response = client.clusters.begin_upgrade(
        resource_group_name="hiloResourcegroup",
        cluster_pool_name="clusterpool1",
        cluster_name="cluster1",
        cluster_upgrade_request={
            "properties": {
                "componentName": "historyserver",
                "targetBuildNumber": "3",
                "targetClusterVersion": "1.0.6",
                "targetOssVersion": "1.16.0",
                "upgradeType": "HotfixUpgrade",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/UpgradeHotfixForCluster.json
if __name__ == "__main__":
    main()
