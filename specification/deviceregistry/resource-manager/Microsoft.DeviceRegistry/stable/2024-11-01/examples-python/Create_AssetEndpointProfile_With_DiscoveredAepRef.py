from azure.identity import DefaultAzureCredential

from azure.mgmt.deviceregistry import DeviceRegistryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceregistry
# USAGE
    python create_asset_endpoint_profile_with_discovered_aep_ref.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DeviceRegistryMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.asset_endpoint_profiles.begin_create_or_replace(
        resource_group_name="myResourceGroup",
        asset_endpoint_profile_name="my-assetendpointprofile",
        resource={
            "extendedLocation": {
                "name": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1",
                "type": "CustomLocation",
            },
            "location": "West Europe",
            "properties": {
                "authentication": {"method": "Anonymous"},
                "discoveredAssetEndpointProfileRef": "discoveredAssetEndpointProfile1",
                "endpointProfileType": "myEndpointProfileType",
                "targetAddress": "https://www.example.com/myTargetAddress",
            },
            "tags": {"site": "building-1"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-11-01/Create_AssetEndpointProfile_With_DiscoveredAepRef.json
if __name__ == "__main__":
    main()
