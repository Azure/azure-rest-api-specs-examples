const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of usage per SKU family for the specified subscription in the specified region.
 *
 * @summary Returns list of usage per SKU family for the specified subscription in the specified region.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Usages/getUsages.json
 */
async function listUsages() {
  const subscriptionId =
    process.env["LABSERVICES_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "eastus2";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
