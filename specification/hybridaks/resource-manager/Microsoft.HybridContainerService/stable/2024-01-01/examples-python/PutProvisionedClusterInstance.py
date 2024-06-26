from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridcontainerservice import HybridContainerServiceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcontainerservice
# USAGE
    python put_provisioned_cluster_instance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridContainerServiceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.provisioned_cluster_instances.begin_create_or_update(
        connected_cluster_resource_uri="subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster",
        provisioned_cluster_instance={
            "extendedLocation": {
                "name": "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation",
                "type": "CustomLocation",
            },
            "properties": {
                "agentPoolProfiles": [
                    {
                        "count": 1,
                        "name": "default-nodepool-1",
                        "nodeLabels": {"env": "dev", "goal": "test"},
                        "nodeTaints": ["env=prod:NoSchedule", "sku=gpu:NoSchedule"],
                        "osType": "Linux",
                        "vmSize": "Standard_A4_v2",
                    }
                ],
                "cloudProviderProfile": {
                    "infraNetworkProfile": {
                        "vnetSubnetIds": [
                            "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.AzureStackHCI/logicalNetworks/test-vnet-static"
                        ]
                    }
                },
                "clusterVMAccessProfile": {"authorizedIPRanges": "127.0.0.1,127.0.0.2"},
                "controlPlane": {"count": 1, "vmSize": "Standard_A4_v2"},
                "kubernetesVersion": "v1.20.5",
                "licenseProfile": {"azureHybridBenefit": "NotApplicable"},
                "linuxProfile": {
                    "ssh": {"publicKeys": [{"keyData": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."}]}
                },
                "networkProfile": {"networkPolicy": "calico", "podCidr": "10.244.0.0/16"},
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/PutProvisionedClusterInstance.json
if __name__ == "__main__":
    main()
