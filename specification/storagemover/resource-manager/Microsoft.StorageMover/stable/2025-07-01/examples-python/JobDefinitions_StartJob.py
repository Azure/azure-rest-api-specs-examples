from azure.identity import DefaultAzureCredential

from azure.mgmt.storagemover import StorageMoverMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagemover
# USAGE
    python job_definitions_start_job.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageMoverMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.job_definitions.start_job(
        resource_group_name="examples-rg",
        storage_mover_name="examples-storageMoverName",
        project_name="examples-projectName",
        job_definition_name="examples-jobDefinitionName",
    )
    print(response)


# x-ms-original-file: 2025-07-01/JobDefinitions_StartJob.json
if __name__ == "__main__":
    main()
