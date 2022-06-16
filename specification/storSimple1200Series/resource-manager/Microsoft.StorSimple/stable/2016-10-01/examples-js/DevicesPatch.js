const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the device.
 *
 * @summary Patches the device.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesPatch.json
 */
async function devicesPatch() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-UGU4PITWNI";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const devicePatch = {
    deviceDescription: "NewDescription8/14/2018 2:30:34 PM",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.devices.beginPatchAndWait(
    deviceName,
    resourceGroupName,
    managerName,
    devicePatch
  );
  console.log(result);
}

devicesPatch().catch(console.error);
