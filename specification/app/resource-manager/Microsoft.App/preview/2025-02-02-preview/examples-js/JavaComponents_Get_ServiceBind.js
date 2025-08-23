const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a Java Component.
 *
 * @summary Get a Java Component.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/JavaComponents_Get_ServiceBind.json
 */
async function getJavaComponentWithServiceBinds() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "myjavacomponent";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.javaComponents.get(resourceGroupName, environmentName, name);
  console.log(result);
}
