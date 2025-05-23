from azure.identity import DefaultAzureCredential

from azure.mgmt.containerinstance import ContainerInstanceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerinstance
# USAGE
    python container_group_profile_create_or_update_extensions.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerInstanceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.container_group_profiles.create_or_update(
        resource_group_name="demo",
        container_group_profile_name="demo1",
        container_group_profile={
            "location": "eastus2",
            "properties": {
                "containers": [
                    {
                        "name": "demo1",
                        "properties": {
                            "command": [],
                            "environmentVariables": [],
                            "image": "nginx",
                            "ports": [{"port": 80}],
                            "resources": {"requests": {"cpu": 1, "memoryInGB": 1.5}},
                        },
                    }
                ],
                "extensions": [
                    {
                        "name": "kube-proxy",
                        "properties": {
                            "extensionType": "kube-proxy",
                            "protectedSettings": {"kubeConfig": "<kubeconfig encoded string>"},
                            "settings": {"clusterCidr": "10.240.0.0/16", "kubeVersion": "v1.9.10"},
                            "version": "1.0",
                        },
                    },
                    {
                        "name": "vk-realtime-metrics",
                        "properties": {"extensionType": "realtime-metrics", "version": "1.0"},
                    },
                ],
                "imageRegistryCredentials": [],
                "ipAddress": {"ports": [{"port": 80, "protocol": "TCP"}], "type": "Private"},
                "osType": "Linux",
            },
            "zones": ["1"],
        },
    )
    print(response)


# x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfileCreateOrUpdate_Extensions.json
if __name__ == "__main__":
    main()
