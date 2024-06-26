from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python kusto_pool_private_link_resources_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="7a587823-959d-4ad0-85bd-cf2a7cef436a",
    )

    response = client.kusto_pool_private_link_resources.list(
        resource_group_name="DP-900",
        workspace_name="synapse-ws-ebi-data",
        kusto_pool_name="dataexplorerpool900",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrivateLinkResourcesList.json
if __name__ == "__main__":
    main()
