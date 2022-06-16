const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clear the alerts.
 *
 * @summary Clear the alerts.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/AlertsClear.json
 */
async function alertsClear() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    alerts: [
      "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/alerts/308b5bd2-824b-436f-840e-44bde075ef33",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.alerts.clear(resourceGroupName, managerName, parameters);
  console.log(result);
}

alertsClear().catch(console.error);
