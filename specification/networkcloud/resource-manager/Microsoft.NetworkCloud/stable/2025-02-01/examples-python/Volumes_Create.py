from azure.identity import DefaultAzureCredential

from azure.mgmt.networkcloud import NetworkCloudMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-networkcloud
# USAGE
    python volumes_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkCloudMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="123e4567-e89b-12d3-a456-426655440000",
    )

    response = client.volumes.begin_create_or_update(
        resource_group_name="resourceGroupName",
        volume_name="volumeName",
        volume_parameters={
            "extendedLocation": {
                "name": "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName",
                "type": "CustomLocation",
            },
            "location": "location",
            "properties": {"sizeMiB": 10000},
            "tags": {"key1": "myvalue1", "key2": "myvalue2"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Volumes_Create.json
if __name__ == "__main__":
    main()
