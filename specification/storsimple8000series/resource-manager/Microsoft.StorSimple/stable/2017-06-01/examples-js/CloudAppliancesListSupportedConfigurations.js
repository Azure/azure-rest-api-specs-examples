const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function cloudAppliancesListSupportedConfigurations() {
  const subscriptionId = "d3ebfe71-b7a9-4c57-92b9-68a2afde4de5";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cloudAppliances.listSupportedConfigurations(
    resourceGroupName,
    managerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

cloudAppliancesListSupportedConfigurations().catch(console.error);
