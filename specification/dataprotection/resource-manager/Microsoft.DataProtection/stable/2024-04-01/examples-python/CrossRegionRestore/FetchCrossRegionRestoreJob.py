from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.dataprotection import DataProtectionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dataprotection
# USAGE
    python fetch_cross_region_restore_job.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataProtectionMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="62b829ee-7936-40c9-a1c9-47a93f9f3965",
    )

    response = client.fetch_cross_region_restore_job.get(
        resource_group_name="BugBash1",
        location="west us",
        parameters={
            "jobId": "3c60cb49-63e8-4b21-b9bd-26277b3fdfae",
            "sourceBackupVaultId": "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/BugBash1/providers/Microsoft.DataProtection/backupVaults/BugBashVaultForCCYv11",
            "sourceRegion": "east us",
        },
    )
    print(response)


# x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/CrossRegionRestore/FetchCrossRegionRestoreJob.json
if __name__ == "__main__":
    main()
