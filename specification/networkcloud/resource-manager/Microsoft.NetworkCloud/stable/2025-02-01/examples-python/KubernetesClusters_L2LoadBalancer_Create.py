from azure.identity import DefaultAzureCredential

from azure.mgmt.networkcloud import NetworkCloudMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkcloud
# USAGE
    python kubernetes_clusters_l2_load_balancer_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkCloudMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="123e4567-e89b-12d3-a456-426655440000",
    )

    response = client.kubernetes_clusters.begin_create_or_update(
        resource_group_name="resourceGroupName",
        kubernetes_cluster_name="kubernetesClusterName",
        kubernetes_cluster_parameters={
            "extendedLocation": {
                "name": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName",
                "type": "CustomLocation",
            },
            "location": "location",
            "properties": {
                "aadConfiguration": {"adminGroupObjectIds": ["ffffffff-ffff-ffff-ffff-ffffffffffff"]},
                "administratorConfiguration": {
                    "adminUsername": "azure",
                    "sshPublicKeys": [
                        {
                            "keyData": "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"
                        }
                    ],
                },
                "controlPlaneNodeConfiguration": {
                    "administratorConfiguration": {
                        "adminUsername": "azure",
                        "sshPublicKeys": [
                            {
                                "keyData": "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"
                            }
                        ],
                    },
                    "availabilityZones": ["1", "2", "3"],
                    "count": 3,
                    "vmSkuName": "NC_G6_28_v1",
                },
                "initialAgentPoolConfigurations": [
                    {
                        "administratorConfiguration": {
                            "adminUsername": "azure",
                            "sshPublicKeys": [
                                {
                                    "keyData": "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"
                                }
                            ],
                        },
                        "agentOptions": {"hugepagesCount": 96, "hugepagesSize": "1G"},
                        "attachedNetworkConfiguration": {
                            "l2Networks": [
                                {
                                    "networkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName",
                                    "pluginType": "DPDK",
                                }
                            ],
                            "l3Networks": [
                                {
                                    "ipamEnabled": "False",
                                    "networkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName",
                                    "pluginType": "SRIOV",
                                }
                            ],
                            "trunkedNetworks": [
                                {
                                    "networkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName",
                                    "pluginType": "MACVLAN",
                                }
                            ],
                        },
                        "availabilityZones": ["1", "2", "3"],
                        "count": 3,
                        "labels": [{"key": "kubernetes.label", "value": "true"}],
                        "mode": "System",
                        "name": "SystemPool-1",
                        "taints": [{"key": "kubernetes.taint", "value": "true:NoSchedule"}],
                        "upgradeSettings": {"maxSurge": "1"},
                        "vmSkuName": "NC_P46_224_v1",
                    }
                ],
                "kubernetesVersion": "1.XX.Y",
                "managedResourceGroupConfiguration": {"location": "East US", "name": "my-managed-rg"},
                "networkConfiguration": {
                    "attachedNetworkConfiguration": {
                        "l2Networks": [
                            {
                                "networkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName",
                                "pluginType": "DPDK",
                            }
                        ],
                        "l3Networks": [
                            {
                                "ipamEnabled": "False",
                                "networkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName",
                                "pluginType": "SRIOV",
                            }
                        ],
                        "trunkedNetworks": [
                            {
                                "networkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName",
                                "pluginType": "MACVLAN",
                            }
                        ],
                    },
                    "cloudServicesNetworkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName",
                    "cniNetworkId": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName",
                    "dnsServiceIp": "198.51.101.2",
                    "l2ServiceLoadBalancerConfiguration": {
                        "ipAddressPools": [
                            {"addresses": ["198.51.102.2-198.51.102.254"], "autoAssign": "True", "name": "pool1"}
                        ]
                    },
                    "podCidrs": ["198.51.100.0/24"],
                    "serviceCidrs": ["198.51.101.0/24"],
                },
            },
            "tags": {"key1": "myvalue1", "key2": "myvalue2"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/KubernetesClusters_L2LoadBalancer_Create.json
if __name__ == "__main__":
    main()
