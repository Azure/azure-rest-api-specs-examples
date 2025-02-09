from azure.identity import DefaultAzureCredential

from azure.mgmt.monitor import MonitorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-monitor
# USAGE
    python private_link_scope_operation_statuses.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MonitorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="613192d7-503f-477a-9cfe-4efc3ee2bd60",
    )

    response = client.private_link_scope_operation_status.get(
        async_operation_id="713192d7-503f-477a-9cfe-4efc3ee2bd11",
        resource_group_name="MyResourceGroup",
    )
    print(response)


# x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2019-10-17-preview/examples/privateLinkScopeOperationStatuses.json
if __name__ == "__main__":
    main()
