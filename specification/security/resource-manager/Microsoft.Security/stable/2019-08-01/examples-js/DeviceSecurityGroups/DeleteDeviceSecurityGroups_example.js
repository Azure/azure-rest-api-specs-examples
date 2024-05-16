const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to User this method to deletes the device security group.
 *
 * @summary User this method to deletes the device security group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/DeleteDeviceSecurityGroups_example.json
 */
async function deleteADeviceSecurityGroupForTheSpecifiedIoTHubResource() {
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub";
  const deviceSecurityGroupName = "samplesecuritygroup";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.deviceSecurityGroups.delete(resourceId, deviceSecurityGroupName);
  console.log(result);
}
