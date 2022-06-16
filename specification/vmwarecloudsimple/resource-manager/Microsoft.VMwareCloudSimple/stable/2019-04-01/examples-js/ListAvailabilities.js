const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of available resources in region
 *
 * @summary Returns list of available resources in region
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListAvailabilities.json
 */
async function listAvailabilities() {
  const subscriptionId = "{subscription-id}";
  const regionId = "westus2";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.skusAvailability.list(regionId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAvailabilities().catch(console.error);
