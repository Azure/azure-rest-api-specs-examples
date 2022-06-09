```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to creates or updates the device security group on a specified IoT Hub resource.
 *
 * @summary Use this method to creates or updates the device security group on a specified IoT Hub resource.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/PutDeviceSecurityGroups_example.json
 */
async function createOrUpdateADeviceSecurityGroupForTheSpecifiedIoTHubResource() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub";
  const deviceSecurityGroupName = "samplesecuritygroup";
  const deviceSecurityGroup = {
    timeWindowRules: [
      {
        isEnabled: true,
        maxThreshold: 30,
        minThreshold: 0,
        ruleType: "ActiveConnectionsNotInAllowedRange",
        timeWindowSize: "PT05M",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.deviceSecurityGroups.createOrUpdate(
    resourceId,
    deviceSecurityGroupName,
    deviceSecurityGroup
  );
  console.log(result);
}

createOrUpdateADeviceSecurityGroupForTheSpecifiedIoTHubResource().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
