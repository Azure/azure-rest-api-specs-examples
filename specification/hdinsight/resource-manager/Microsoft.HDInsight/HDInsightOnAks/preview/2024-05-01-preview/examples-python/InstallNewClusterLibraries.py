from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsightcontainers import HDInsightContainersMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsightcontainers
# USAGE
    python install_new_cluster_libraries.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HDInsightContainersMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.cluster_libraries.begin_manage_libraries(
        resource_group_name="hiloResourceGroup",
        cluster_pool_name="clusterPool",
        cluster_name="cluster",
        operation={
            "properties": {
                "action": "Install",
                "libraries": [
                    {
                        "properties": {
                            "name": "requests",
                            "remarks": "PyPi packages.",
                            "type": "pypi",
                            "version": "2.31.0",
                        }
                    },
                    {
                        "properties": {
                            "groupId": "org.apache.flink",
                            "name": "flink-connector-kafka",
                            "remarks": "Maven packages.",
                            "type": "maven",
                            "version": "3.0.2-1.18",
                        }
                    },
                ],
            }
        },
    ).result()


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/InstallNewClusterLibraries.json
if __name__ == "__main__":
    main()
