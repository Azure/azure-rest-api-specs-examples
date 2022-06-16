const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Dapr Components for a managed environment.
 *
 * @summary Get the Dapr Components for a managed environment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/DaprComponents_List.json
 */
async function listDaprComponents() {
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = "examplerg";
  const environmentName = "myenvironment";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.daprComponents.list(resourceGroupName, environmentName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDaprComponents().catch(console.error);
