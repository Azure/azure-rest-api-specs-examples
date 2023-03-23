const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get the network endpoints of all outbound dependencies of an App Service Environment.
 *
 * @summary Description for Get the network endpoints of all outbound dependencies of an App Service Environment.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/GetOutboundNetworkDependenciesEndpoints.json
 */
async function getTheNetworkEndpointsOfAllOutboundDependenciesOfAnAppServiceEnvironment() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName =
    process.env["APPSERVICE_RESOURCE_GROUP"] || "Sample-WestUSResourceGroup";
  const name = "SampleAse";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appServiceEnvironments.listOutboundNetworkDependenciesEndpoints(
    resourceGroupName,
    name
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
