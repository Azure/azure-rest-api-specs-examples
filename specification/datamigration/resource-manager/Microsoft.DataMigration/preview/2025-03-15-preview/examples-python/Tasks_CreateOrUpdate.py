from azure.identity import DefaultAzureCredential

from azure.mgmt.datamigration import DataMigrationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datamigration
# USAGE
    python tasks_create_or_update.py

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

    response = client.tasks.create_or_update(
        group_name="DmsSdkRg",
        service_name="DmsSdkService",
        project_name="DmsSdkProject",
        task_name="DmsSdkTask",
        parameters={
            "properties": {
                "input": {
                    "targetConnectionInfo": {
                        "authentication": "SqlAuthentication",
                        "dataSource": "ssma-test-server.database.windows.net",
                        "encryptConnection": True,
                        "password": "testpassword",
                        "trustServerCertificate": True,
                        "type": "SqlConnectionInfo",
                        "userName": "testuser",
                    }
                },
                "taskType": "ConnectToTarget.SqlDb",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/Tasks_CreateOrUpdate.json
if __name__ == "__main__":
    main()
