from azure.identity import DefaultAzureCredential

from azure.mgmt.hdinsight import HDInsightManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hdinsight
# USAGE
    python create_hd_insight_cluster_with_tls12.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HDInsightManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.clusters.begin_create(
        resource_group_name="rg1",
        cluster_name="cluster1",
        parameters={
            "properties": {
                "clusterDefinition": {
                    "configurations": {
                        "gateway": {
                            "restAuthCredential.isEnabled": True,
                            "restAuthCredential.password": "**********",
                            "restAuthCredential.username": "admin",
                        }
                    },
                    "kind": "Hadoop",
                },
                "clusterVersion": "3.6",
                "computeProfile": {
                    "roles": [
                        {
                            "hardwareProfile": {"vmSize": "Large"},
                            "name": "headnode",
                            "osProfile": {
                                "linuxOperatingSystemProfile": {"password": "**********", "username": "sshuser"}
                            },
                            "targetInstanceCount": 2,
                        },
                        {
                            "hardwareProfile": {"vmSize": "Large"},
                            "name": "workernode",
                            "osProfile": {
                                "linuxOperatingSystemProfile": {"password": "**********", "username": "sshuser"}
                            },
                            "targetInstanceCount": 3,
                        },
                        {
                            "hardwareProfile": {"vmSize": "Small"},
                            "name": "zookeepernode",
                            "osProfile": {
                                "linuxOperatingSystemProfile": {"password": "**********", "username": "sshuser"}
                            },
                            "targetInstanceCount": 3,
                        },
                    ]
                },
                "minSupportedTlsVersion": "1.2",
                "osType": "Linux",
                "storageProfile": {
                    "storageaccounts": [
                        {
                            "container": "default8525",
                            "enableSecureChannel": True,
                            "isDefault": True,
                            "key": "storagekey",
                            "name": "mystorage.blob.core.windows.net",
                        }
                    ]
                },
                "tier": "Standard",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateHDInsightClusterWithTLS12.json
if __name__ == "__main__":
    main()
