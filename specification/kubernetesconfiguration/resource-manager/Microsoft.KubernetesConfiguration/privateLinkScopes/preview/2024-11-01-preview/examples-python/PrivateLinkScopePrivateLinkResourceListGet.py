from azure.identity import DefaultAzureCredential

from azure.mgmt.kubernetesconfiguration.privatelinkscopes import KubernetesConfigurationPrivateLinkScopesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kubernetesconfiguration-privatelinkscopes
# USAGE
    python private_link_scope_private_link_resource_list_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KubernetesConfigurationPrivateLinkScopesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.private_link_resources.list_by_private_link_scope(
        resource_group_name="myResourceGroup",
        scope_name="myPrivateLinkScope",
    )
    print(response)


# x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateLinkScopePrivateLinkResourceListGet.json
if __name__ == "__main__":
    main()
