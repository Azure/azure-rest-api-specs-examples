from azure.identity import DefaultAzureCredential

from azure.mgmt.kubernetesconfiguration.privatelinkscopes import KubernetesConfigurationPrivateLinkScopesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kubernetesconfiguration-privatelinkscopes
# USAGE
    python private_link_scopes_list_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KubernetesConfigurationPrivateLinkScopesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="86dc51d3-92ed-4d7e-947a-775ea79b4919",
    )

    response = client.private_link_scopes.list_by_resource_group(
        resource_group_name="my-resource-group",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateLinkScopesListByResourceGroup.json
if __name__ == "__main__":
    main()
