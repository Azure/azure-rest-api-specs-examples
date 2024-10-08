from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsightcontainers import HDInsightContainersMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsightcontainers
# USAGE
    python patch_ranger_cluster.py

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

    response = client.clusters.begin_update(
        resource_group_name="hiloResourcegroup",
        cluster_pool_name="clusterpool1",
        cluster_name="cluster1",
        cluster_patch_request={
            "properties": {
                "clusterProfile": {
                    "rangerProfile": {
                        "rangerAdmin": {
                            "admins": ["testuser1@contoso.com", "testuser2@contoso.com"],
                            "database": {
                                "host": "testsqlserver.database.windows.net",
                                "name": "testdb",
                                "passwordSecretRef": "https://testkv.vault.azure.net/secrets/mysecret/5df6584d9c25418c8d900240aa6c3452",
                                "username": "admin",
                            },
                        },
                        "rangerAudit": {"storageAccount": "https://teststorage.blob.core.windows.net/testblob"},
                        "rangerUsersync": {
                            "enabled": True,
                            "groups": ["0a53828f-36c9-44c3-be3d-99a7fce977ad", "13be6971-79db-4f33-9d41-b25589ca25ac"],
                            "mode": "automatic",
                            "users": ["testuser1@contoso.com", "testuser2@contoso.com"],
                        },
                    }
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/PatchRangerCluster.json
if __name__ == "__main__":
    main()
