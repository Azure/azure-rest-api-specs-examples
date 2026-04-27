from azure.identity import DefaultAzureCredential

from azure.mgmt.kubernetesconfiguration.extensions import KubernetesConfigurationExtensionsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kubernetesconfiguration-extensions
# USAGE
    python list_extensions.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KubernetesConfigurationExtensionsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.extensions.list(
        resource_group_name="rg1",
        cluster_rp="Microsoft.Kubernetes",
        cluster_resource_name="connectedClusters",
        cluster_name="clusterName1",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2025-03-01/ListExtensions.json
if __name__ == "__main__":
    main()
