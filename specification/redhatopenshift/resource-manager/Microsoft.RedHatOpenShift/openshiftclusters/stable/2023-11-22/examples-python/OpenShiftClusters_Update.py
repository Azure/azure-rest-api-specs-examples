from azure.identity import DefaultAzureCredential

from azure.mgmt.redhatopenshift import AzureRedHatOpenShiftClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-redhatopenshift
# USAGE
    python open_shift_clusters_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureRedHatOpenShiftClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.open_shift_clusters.begin_update(
        resource_group_name="resourceGroup",
        resource_name="resourceName",
        parameters={
            "properties": {
                "apiserverProfile": {"visibility": "Public"},
                "clusterProfile": {
                    "domain": "cluster.location.aroapp.io",
                    "fipsValidatedModules": "Enabled",
                    "pullSecret": '{"auths":{"registry.connect.redhat.com":{"auth":""},"registry.redhat.io":{"auth":""}}}',
                    "resourceGroupId": "/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup",
                },
                "consoleProfile": {},
                "ingressProfiles": [{"name": "default", "visibility": "Public"}],
                "masterProfile": {
                    "encryptionAtHost": "Enabled",
                    "subnetId": "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master",
                    "vmSize": "Standard_D8s_v3",
                },
                "networkProfile": {
                    "loadBalancerProfile": {"managedOutboundIps": {"count": 1}},
                    "podCidr": "10.128.0.0/14",
                    "preconfiguredNSG": "Disabled",
                    "serviceCidr": "172.30.0.0/16",
                },
                "servicePrincipalProfile": {"clientId": "clientId", "clientSecret": "clientSecret"},
                "workerProfiles": [
                    {
                        "count": 3,
                        "diskSizeGB": 128,
                        "name": "worker",
                        "subnetId": "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker",
                        "vmSize": "Standard_D2s_v3",
                    }
                ],
            },
            "tags": {"key": "value"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/OpenShiftClusters_Update.json
if __name__ == "__main__":
    main()
