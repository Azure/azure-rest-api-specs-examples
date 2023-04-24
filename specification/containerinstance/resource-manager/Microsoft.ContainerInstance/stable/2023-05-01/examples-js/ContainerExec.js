const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Executes a command for a specific container instance in a specified resource group and container group.
 *
 * @summary Executes a command for a specific container instance in a specified resource group and container group.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerExec.json
 */
async function containerExec() {
  const subscriptionId = process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupName = "demo1";
  const containerName = "container1";
  const containerExecRequest = {
    command: "/bin/bash",
    terminalSize: { cols: 12, rows: 12 },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containers.executeCommand(
    resourceGroupName,
    containerGroupName,
    containerName,
    containerExecRequest
  );
  console.log(result);
}
