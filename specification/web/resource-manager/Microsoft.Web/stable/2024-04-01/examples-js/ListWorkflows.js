const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List the workflows for a web site, or a deployment slot.
 *
 * @summary List the workflows for a web site, or a deployment slot.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/ListWorkflows.json
 */
async function listTheWorkflows() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "testsite2";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.webApps.listWorkflows(resourceGroupName, name)) {
    resArray.push(item);
  }
  console.log(resArray);
}
