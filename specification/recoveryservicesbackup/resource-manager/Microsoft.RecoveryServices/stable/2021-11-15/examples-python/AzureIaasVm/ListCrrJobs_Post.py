from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.passivestamp import RecoveryServicesBackupPassiveClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python list_crr_jobs_post.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupPassiveClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.backup_crr_jobs.list(
        azure_region="southeastasia",
        parameters={
            "jobName": "02585cc9-d7f4-4b46-860c-14c048cce178",
            "resourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Compute/VirtualMachines/testVm",
        },
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-11-15/examples/AzureIaasVm/ListCrrJobs_Post.json
if __name__ == "__main__":
    main()
