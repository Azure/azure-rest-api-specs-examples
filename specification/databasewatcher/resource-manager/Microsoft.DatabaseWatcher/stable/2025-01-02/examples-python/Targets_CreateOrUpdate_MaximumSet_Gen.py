from azure.identity import DefaultAzureCredential

from azure.mgmt.databasewatcher import DatabaseWatcherMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databasewatcher
# USAGE
    python targets_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DatabaseWatcherMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.targets.create_or_update(
        resource_group_name="apiTest-ddat4p",
        watcher_name="databasemo3ej9ih",
        target_name="monitoringh22eed",
        resource={
            "properties": {
                "connectionServerName": "sqlServero1ihe2",
                "sqlDbResourceId": "/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2",
                "targetAuthenticationType": "Aad",
                "targetType": "SqlDb",
            }
        },
    )
    print(response)


# x-ms-original-file: 2025-01-02/Targets_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
