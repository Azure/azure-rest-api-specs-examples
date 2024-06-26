from azure.identity import DefaultAzureCredential
from azure.mgmt.storagemover import StorageMoverMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagemover
# USAGE
    python endpoints_update_nfs_mount.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageMoverMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="60bcfc77-6589-4da2-b7fd-f9ec9322cf95",
    )

    response = client.endpoints.update(
        resource_group_name="examples-rg",
        storage_mover_name="examples-storageMoverName",
        endpoint_name="examples-endpointName",
        endpoint={"properties": {"description": "Updated Endpoint Description", "endpointType": "NfsMount"}},
    )
    print(response)


# x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Endpoints_Update_NfsMount.json
if __name__ == "__main__":
    main()
