from azure.identity import DefaultAzureCredential

from azure.mgmt.datamigration import DataMigrationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datamigration
# USAGE
    python tasks_command.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataMigrationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="fc04246f-04c5-437e-ac5e-206a19e7193f",
    )

    response = client.tasks.command(
        group_name="DmsSdkRg",
        service_name="DmsSdkService",
        project_name="DmsSdkProject",
        task_name="DmsSdkTask",
        parameters={"commandType": "Migrate.Sync.Complete.Database", "input": {"databaseName": "TestDatabase"}},
    )
    print(response)


# x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/Tasks_Command.json
if __name__ == "__main__":
    main()
