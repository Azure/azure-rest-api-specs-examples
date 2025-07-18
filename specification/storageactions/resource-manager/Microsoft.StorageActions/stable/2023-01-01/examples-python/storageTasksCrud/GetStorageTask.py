from azure.identity import DefaultAzureCredential

from azure.mgmt.storageactions import StorageActionsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storageactions
# USAGE
    python get_storage_task.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageActionsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.storage_tasks.get(
        resource_group_name="res4228",
        storage_task_name="mytask1",
    )
    print(response)


# x-ms-original-file: 2023-01-01/storageTasksCrud/GetStorageTask.json
if __name__ == "__main__":
    main()
