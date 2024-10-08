from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsightcontainers import HDInsightContainersMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsightcontainers
# USAGE
    python create_spark_cluster.py

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

    response = client.clusters.begin_create(
        resource_group_name="hiloResourcegroup",
        cluster_pool_name="clusterpool1",
        cluster_name="cluster1",
        hd_insight_cluster={
            "location": "West US 2",
            "properties": {
                "clusterProfile": {
                    "authorizationProfile": {"userIds": ["testuser1", "testuser2"]},
                    "clusterVersion": "0.0.1",
                    "managedIdentityProfile": {
                        "identityList": [
                            {
                                "clientId": "de91f1d8-767f-460a-ac11-3cf103f74b34",
                                "objectId": "40491351-c240-4042-91e0-f644a1d2b441",
                                "resourceId": "/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi",
                                "type": "cluster",
                            }
                        ]
                    },
                    "ossVersion": "2.2.3",
                    "serviceConfigsProfiles": [
                        {
                            "configs": [
                                {
                                    "component": "spark-config",
                                    "files": [
                                        {
                                            "fileName": "spark-defaults.conf",
                                            "values": {"spark.eventLog.enabled": "true"},
                                        }
                                    ],
                                }
                            ],
                            "serviceName": "spark-service",
                        },
                        {
                            "configs": [
                                {
                                    "component": "yarn-config",
                                    "files": [
                                        {
                                            "fileName": "core-site.xml",
                                            "values": {
                                                "fs.defaultFS": "wasb://testcontainer@teststorage.dfs.core.windows.net/",
                                                "storage.container": "testcontainer",
                                                "storage.key": "test key",
                                                "storage.name": "teststorage",
                                                "storage.protocol": "wasb",
                                            },
                                        },
                                        {"fileName": "yarn-site.xml", "values": {"yarn.webapp.ui2.enable": "false"}},
                                    ],
                                }
                            ],
                            "serviceName": "yarn-service",
                        },
                    ],
                    "sparkProfile": {},
                    "sshProfile": {"count": 2, "vmSize": "Standard_D3_v2"},
                },
                "clusterType": "spark",
                "computeProfile": {
                    "availabilityZones": ["1", "2", "3"],
                    "nodes": [{"count": 4, "type": "worker", "vmSize": "Standard_D3_v2"}],
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/CreateSparkCluster.json
if __name__ == "__main__":
    main()
