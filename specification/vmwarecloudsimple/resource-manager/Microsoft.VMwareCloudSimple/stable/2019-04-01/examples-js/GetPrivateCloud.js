const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns private cloud by its name
 *
 * @summary Returns private cloud by its name
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetPrivateCloud.json
 */
async function getPrivateCloud() {
  const subscriptionId = "{subscription-id}";
  const pcName = "myPrivateCloud";
  const regionId = "westus2";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.privateClouds.get(pcName, regionId);
  console.log(result);
}

getPrivateCloud().catch(console.error);
