from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridcontainerservice import HybridContainerServiceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcontainerservice
# USAGE
    python put_provisioned_cluster.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridContainerServiceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b",
    )

    response = client.provisioned_clusters.begin_create_or_update(
        resource_group_name="test-arcappliance-resgrp",
        resource_name="test-hybridakscluster",
        provisioned_clusters={
            "extendedLocation": {
                "name": "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation",
                "type": "CustomLocation",
            },
            "location": "westus",
            "properties": {
                "agentPoolProfiles": [
                    {"count": 1, "name": "default-nodepool-1", "osType": "Linux", "vmSize": "Standard_A4_v2"}
                ],
                "cloudProviderProfile": {
                    "infraNetworkProfile": {
                        "vnetSubnetIds": [
                            "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/virtualNetworks/test-vnet-static"
                        ]
                    },
                    "infraStorageProfile": {
                        "storageSpaceIds": [
                            "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/storageSpaces/test-storage"
                        ]
                    },
                },
                "controlPlane": {
                    "count": 1,
                    "linuxProfile": {
                        "ssh": {"publicKeys": [{"keyData": "ssh-rsa AAAAB1NzaC1yc2EAAAADAQABAAACAQCY......"}]}
                    },
                    "osType": "Linux",
                    "vmSize": "Standard_A4_v2",
                },
                "kubernetesVersion": "v1.20.5",
                "linuxProfile": {
                    "ssh": {"publicKeys": [{"keyData": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."}]}
                },
                "networkProfile": {
                    "loadBalancerProfile": {
                        "count": 1,
                        "linuxProfile": {
                            "ssh": {"publicKeys": [{"keyData": "ssh-rsa AAAAB2NzaC1yc2EAAAADAQABAAACAQCY......"}]}
                        },
                        "osType": "Linux",
                        "vmSize": "Standard_K8S3_v1",
                    },
                    "loadBalancerSku": "unstacked-haproxy",
                    "networkPolicy": "calico",
                    "podCidr": "10.244.0.0/16",
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/PutProvisionedCluster.json
if __name__ == "__main__":
    main()
