from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.storagecache import StorageCacheManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagecache
# USAGE
    python import_job_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageCacheManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.import_jobs.begin_update(
        resource_group_name="scgroup",
        aml_filesystem_name="fs1",
        import_job_name="job1",
        import_job={"tags": {"Dept": "ContosoAds"}},
    ).result()
    print(response)


# x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/importJob_Update.json
if __name__ == "__main__":
    main()
