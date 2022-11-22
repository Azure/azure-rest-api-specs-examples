from azure.identity import DefaultAzureCredential
from azure.mgmt.datamigration import DataMigrationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datamigration
# USAGE
    python tasks_delete.py

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

    response = client.tasks.delete(
        group_name="DmsSdkRg",
        service_name="DmsSdkService",
        project_name="DmsSdkProject",
        task_name="DmsSdkTask",
    )
    print(response)


# x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/Tasks_Delete.json
if __name__ == "__main__":
    main()
