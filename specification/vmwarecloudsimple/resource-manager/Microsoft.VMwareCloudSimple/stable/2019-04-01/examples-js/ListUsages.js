const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of usage in region
 *
 * @summary Returns list of usage in region
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListUsages.json
 */
async function listUsages() {
  const subscriptionId = "{subscription-id}";
  const regionId = "westus2";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.list(regionId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listUsages().catch(console.error);
