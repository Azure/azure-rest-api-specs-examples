const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a NamespaceDevice
 *
 * @summary create a NamespaceDevice
 * x-ms-original-file: 2025-10-01/CreateOrReplace_NamespaceDevice.json
 */
async function createOrReplaceNamespaceDevices() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.namespaceDevices.createOrReplace(
    "myResourceGroup",
    "adr-namespace-gbk0925-n01",
    "dev-namespace-gbk0925-n01",
    {
      location: "West Europe",
      properties: {
        endpoints: {
          outbound: {
            assigned: {
              eventGridEndpoint: {
                endpointType: "Microsoft.Devices/IoTHubs",
                address: "https://myeventgridtopic.westeurope-1.eventgrid.azure.net/api/events",
              },
            },
          },
        },
        externalDeviceId: "adr-smart-device3-7a848b15-af47-40a7-8c06-a3f43314d44f",
        enabled: true,
        attributes: {
          deviceType: "sensor",
          deviceOwner: "IT",
          deviceCategory: 16,
        },
      },
    },
  );
  console.log(result);
}
