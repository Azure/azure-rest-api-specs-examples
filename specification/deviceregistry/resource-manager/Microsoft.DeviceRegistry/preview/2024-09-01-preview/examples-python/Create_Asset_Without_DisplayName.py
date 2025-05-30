from azure.identity import DefaultAzureCredential

from azure.mgmt.deviceregistry import DeviceRegistryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-deviceregistry
# USAGE
    python create_asset_without_display_name.py

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
                "assetEndpointProfileRef": "myAssetEndpointProfile",
                "datasets": [
                    {
                        "dataPoints": [
                            {
                                "dataPointConfiguration": '{"publishingInterval":8,"samplingInterval":8,"queueSize":4}',
                                "dataSource": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1",
                                "name": "dataPoint1",
                                "observabilityMode": "Counter",
                            },
                            {
                                "dataPointConfiguration": '{"publishingInterval":4,"samplingInterval":4,"queueSize":7}',
                                "dataSource": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2",
                                "name": "dataPoint2",
                                "observabilityMode": "None",
                            },
                        ],
                        "datasetConfiguration": '{"publishingInterval":10,"samplingInterval":15,"queueSize":20}',
                        "name": "dataset1",
                        "topic": {"path": "/path/dataset1", "retain": "Keep"},
                    }
                ],
                "defaultDatasetsConfiguration": '{"publishingInterval":10,"samplingInterval":15,"queueSize":20}',
                "defaultEventsConfiguration": '{"publishingInterval":10,"samplingInterval":15,"queueSize":20}',
                "defaultTopic": {"path": "/path/defaultTopic", "retain": "Keep"},
                "description": "This is a sample Asset",
                "documentationUri": "https://www.example.com/manual",
                "enabled": True,
                "events": [
                    {
                        "eventConfiguration": '{"publishingInterval":7,"samplingInterval":1,"queueSize":8}',
                        "eventNotifier": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3",
                        "name": "event1",
                        "observabilityMode": "None",
                        "topic": {"path": "/path/event1", "retain": "Keep"},
                    },
                    {
                        "eventConfiguration": '{"publishingInterval":7,"samplingInterval":8,"queueSize":4}',
                        "eventNotifier": "nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4",
                        "name": "event2",
                        "observabilityMode": "Log",
                    },
                ],
                "externalAssetId": "8ZBA6LRHU0A458969",
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


# x-ms-original-file: 2024-09-01-preview/Create_Asset_Without_DisplayName.json
if __name__ == "__main__":
    main()
