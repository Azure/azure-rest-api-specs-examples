const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the manager.
 *
 * @summary Creates or updates the manager.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersCreateOrUpdate.json
 */
async function managersCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hManagerForSDKTest";
  const manager = {
    name: "hManagerForSDKTest",
    cisIntrinsicSettings: { type: "HelsinkiV1" },
    location: "westus",
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.managers.createOrUpdate(resourceGroupName, managerName, manager);
  console.log(result);
}

managersCreateOrUpdate().catch(console.error);
