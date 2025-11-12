const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a NamespaceDevice
 *
 * @summary create a NamespaceDevice
 * x-ms-original-file: 2025-10-01/CreateOrReplace_NamespaceDevice_Edge_Anonymous.json
 */
async function createEdgeEnabledDeviceWithAnonymousHostAuthentication() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.namespaceDevices.createOrReplace(
    "myResourceGroup",
    "adr-namespace-gbk0925-n01",
    "namespace-device-on-edge",
    {
      extendedLocation: {
        type: "CustomLocation",
        name: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1",
      },
      location: "West Europe",
      properties: {
        endpoints: {
          inbound: {
            theOnlyOPCUABroker: {
              address: "opc.tcp://192.168.86.23:51211/UA/SampleServer",
              endpointType: "microsoft.opcua",
              version: "2",
              authentication: { method: "Anonymous" },
            },
          },
        },
        externalDeviceId: "unique-edge-device-identifier",
        enabled: true,
        attributes: {
          deviceType: "dough-maker",
          deviceOwner: "OT",
          deviceCategory: 16,
        },
      },
    },
  );
  console.log(result);
}
