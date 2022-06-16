const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of dedicated cloud services within a subscription
 *
 * @summary Returns list of dedicated cloud services within a subscription
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListDedicatedCloudServices.json
 */
async function listDedicatedCloudServices() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dedicatedCloudServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDedicatedCloudServices().catch(console.error);
