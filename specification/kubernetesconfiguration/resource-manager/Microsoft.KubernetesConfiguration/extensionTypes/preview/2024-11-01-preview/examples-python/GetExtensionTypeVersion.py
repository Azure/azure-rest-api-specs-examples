from azure.identity import DefaultAzureCredential

from azure.mgmt.kubernetesconfiguration.extensiontypes import KubernetesConfigurationExtensionTypesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kubernetesconfiguration-extensiontypes
# USAGE
    python get_extension_type_version.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KubernetesConfigurationExtensionTypesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId1",
    )

    response = client.extension_types.cluster_get_version(
        resource_group_name="rg1",
        cluster_rp="Microsoft.Kubernetes",
        cluster_resource_name="connectedClusters",
        cluster_name="my-cluster",
        extension_type_name="my-extension-type",
        version_number="v1.3.2",
    )
    print(response)


# x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/GetExtensionTypeVersion.json
if __name__ == "__main__":
    main()
