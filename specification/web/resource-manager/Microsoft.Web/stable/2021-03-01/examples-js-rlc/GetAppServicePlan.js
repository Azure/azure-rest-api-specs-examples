const WebSiteManagementClient = require("@azure-rest/arm-appservice").default;
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get an App Service plan.
 *
 * @summary Description for Get an App Service plan.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetAppServicePlan.json
 */
async function getAppServicePlan() {
  const subscriptionId = process.env.SUBSCRIPTION_ID;
  const resourceGroupName = "testrg123";
  const name = "testsf6141";
  const credential = new DefaultAzureCredential();
  const client = WebSiteManagementClient(credential);
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}",
      subscriptionId,
      resourceGroupName,
      name
    )
    .get();
  console.log(result);
}

getAppServicePlan().catch(console.error);
