from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python extended_sql_pool_azure_monitor_auditing_create_min.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.extended_sql_pool_blob_auditing_policies.create_or_update(
        resource_group_name="blobauditingtest-4799",
        workspace_name="blobauditingtest-6440",
        sql_pool_name="testdb",
        parameters={"properties": {"isAzureMonitorTargetEnabled": True, "state": "Enabled"}},
    )
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolAzureMonitorAuditingCreateMin.json
if __name__ == "__main__":
    main()
