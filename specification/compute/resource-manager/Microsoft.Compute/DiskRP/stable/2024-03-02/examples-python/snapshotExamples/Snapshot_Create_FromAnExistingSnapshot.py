from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python snapshot_create_from_an_existing_snapshot.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.snapshots.begin_create_or_update(
        resource_group_name="myResourceGroup",
        snapshot_name="mySnapshot2",
        snapshot={
            "location": "West US",
            "properties": {
                "creationData": {
                    "createOption": "Copy",
                    "sourceResourceId": "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1",
                }
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/snapshotExamples/Snapshot_Create_FromAnExistingSnapshot.json
if __name__ == "__main__":
    main()
