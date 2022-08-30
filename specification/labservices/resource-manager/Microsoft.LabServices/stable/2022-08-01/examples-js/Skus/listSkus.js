const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of Azure Lab Services resource SKUs.
 *
 * @summary Returns a list of Azure Lab Services resource SKUs.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Skus/listSkus.json
 */
async function listSkus() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.skus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSkus().catch(console.error);
