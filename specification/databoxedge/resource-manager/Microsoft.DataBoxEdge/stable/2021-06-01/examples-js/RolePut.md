```javascript
const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a role.
 *
 * @summary Create or update a role.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/RolePut.json
 */
async function rolePut() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const name = "IoTRole1";
  const resourceGroupName = "GroupForEdgeAutomation";
  const role = {
    hostPlatform: "Linux",
    ioTDeviceDetails: {
      authentication: {
        symmetricKey: {
          connectionString: {
            encryptionAlgorithm: "AES256",
            encryptionCertThumbprint: "348586569999244",
            value:
              "Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotDevice;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>",
          },
        },
      },
      deviceId: "iotdevice",
      ioTHostHub: "iothub.azure-devices.net",
    },
    ioTEdgeDeviceDetails: {
      authentication: {
        symmetricKey: {
          connectionString: {
            encryptionAlgorithm: "AES256",
            encryptionCertThumbprint: "1245475856069999244",
            value:
              "Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotEdge;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>",
          },
        },
      },
      deviceId: "iotEdge",
      ioTHostHub: "iothub.azure-devices.net",
    },
    kind: "IOT",
    roleStatus: "Enabled",
    shareMappings: [],
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.roles.beginCreateOrUpdateAndWait(
    deviceName,
    name,
    resourceGroupName,
    role
  );
  console.log(result);
}

rolePut().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-databoxedge_2.0.1/sdk/databoxedge/arm-databoxedge/README.md) on how to add the SDK to your project and authenticate.
