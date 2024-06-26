from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python create_workspace_managed_sql_server_extended_blob_auditing_settings_with_all_parameters.py

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

    response = client.workspace_managed_sql_server_extended_blob_auditing_policies.begin_create_or_update(
        resource_group_name="wsg-7398",
        workspace_name="testWorkspace",
        blob_auditing_policy_name="default",
        parameters={
            "properties": {
                "auditActionsAndGroups": [
                    "SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP",
                    "FAILED_DATABASE_AUTHENTICATION_GROUP",
                    "BATCH_COMPLETED_GROUP",
                ],
                "isAzureMonitorTargetEnabled": True,
                "isStorageSecondaryKeyInUse": False,
                "predicateExpression": "object_name = 'SensitiveData'",
                "retentionDays": 6,
                "state": "Enabled",
                "storageAccountAccessKey": "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
                "storageAccountSubscriptionId": "00000000-1234-0000-5678-000000000000",
                "storageEndpoint": "https://mystorage.blob.core.windows.net",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerExtendedBlobAuditingSettingsWithAllParameters.json
if __name__ == "__main__":
    main()
