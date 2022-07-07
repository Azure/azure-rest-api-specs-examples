const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get usage metrics for a multi-role pool of an App Service Environment.
 *
 * @summary Description for Get usage metrics for a multi-role pool of an App Service Environment.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_ListMultiRoleUsages.json
 */
async function getUsageMetricsForAMultiRolePoolOfAnAppServiceEnvironment() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-rg";
  const name = "test-ase";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appServiceEnvironments.listMultiRoleUsages(
    resourceGroupName,
    name
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getUsageMetricsForAMultiRolePoolOfAnAppServiceEnvironment().catch(console.error);
