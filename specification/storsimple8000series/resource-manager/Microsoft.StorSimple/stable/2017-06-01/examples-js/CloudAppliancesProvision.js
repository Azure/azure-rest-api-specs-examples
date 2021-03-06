const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function cloudAppliancesProvision() {
  const subscriptionId = "d3ebfe71-b7a9-4c57-92b9-68a2afde4de5";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    name: "sca07forsdktest",
    modelNumber: "8020",
    vnetRegion: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.cloudAppliances.beginProvisionAndWait(
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

cloudAppliancesProvision().catch(console.error);
