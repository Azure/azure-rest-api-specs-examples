from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.deviceregistry import DeviceRegistryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceregistry
# USAGE
    python create_asset_without_external_asset_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DeviceRegistryMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.assets.begin_create_or_replace(
        resource_group_name="myResourceGroup",
        asset_name="my-asset",
        resource={
            "extendedLocation": {
                "name": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1",
                "type": "CustomLocation",
            },
            "location": "West Europe",
            "properties": {
                "assetEndpointProfileUri": "https://www.example.com/myAssetEndpointProfile",
                "assetType": "MyAssetType",
                "dataPoints": [
                    {
                        "capabilityId": "dtmi:com:example:Thermostat:__temperature;1",
                        "dataPointConfiguration": '{"publishingInterval":8,"samplingInterval":8,"queueSize":4}',
                        "dataSource": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1",
                        "observabilityMode": "counter",
                    },
                    {
                        "capabilityId": "dtmi:com:example:Thermostat:__pressure;1",
                        "dataPointConfiguration": '{"publishingInterval":4,"samplingInterval":4,"queueSize":7}',
                        "dataSource": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2",
                        "observabilityMode": "none",
                    },
                ],
                "defaultDataPointsConfiguration": '{"publishingInterval":10,"samplingInterval":15,"queueSize":20}',
                "defaultEventsConfiguration": '{"publishingInterval":10,"samplingInterval":15,"queueSize":20}',
                "description": "This is a sample Asset",
                "displayName": "AssetDisplayName",
                "documentationUri": "https://www.example.com/manual",
                "enabled": True,
                "events": [
                    {
                        "capabilityId": "dtmi:com:example:Thermostat:__temperature;1",
                        "eventConfiguration": '{"publishingInterval":7,"samplingInterval":1,"queueSize":8}',
                        "eventNotifier": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3",
                        "observabilityMode": "none",
                    },
                    {
                        "capabilityId": "dtmi:com:example:Thermostat:__pressure;1",
                        "eventConfiguration": '{"publishingInterval":7,"samplingInterval":8,"queueSize":4}',
                        "eventNotifier": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4",
                        "observabilityMode": "log",
                    },
                ],
                "hardwareRevision": "1.0",
                "manufacturer": "Contoso",
                "manufacturerUri": "https://www.contoso.com/manufacturerUri",
                "model": "ContosoModel",
                "productCode": "SA34VDG",
                "serialNumber": "64-103816-519918-8",
                "softwareRevision": "2.0",
            },
            "tags": {"site": "building-1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Create_Asset_Without_ExternalAssetId.json
if __name__ == "__main__":
    main()
