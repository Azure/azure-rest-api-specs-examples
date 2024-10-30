const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a .NET Component.
 *
 * @summary Get a .NET Component.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/DotNetComponents_Get_ServiceBind.json
 */
async function getNetComponentWithServiceBinds() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "mydotnetcomponent";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.dotNetComponents.get(resourceGroupName, environmentName, name);
  console.log(result);
}
