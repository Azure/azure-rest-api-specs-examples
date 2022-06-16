const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Changes the power state of the controller.
 *
 * @summary Changes the power state of the controller.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/HardwareComponentGroupsChangeControllerPowerState.json
 */
async function hardwareComponentGroupsChangeControllerPowerState() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const hardwareComponentGroupName = "Controller0Components";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    action: "Start",
    activeController: "Controller0",
    controller0State: "Ok",
    controller1State: "NotPresent",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.hardwareComponentGroups.beginChangeControllerPowerStateAndWait(
    deviceName,
    hardwareComponentGroupName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

hardwareComponentGroupsChangeControllerPowerState().catch(console.error);
