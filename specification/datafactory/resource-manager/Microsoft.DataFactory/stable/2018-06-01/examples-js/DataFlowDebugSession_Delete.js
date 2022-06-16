const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a data flow debug session.
 *
 * @summary Deletes a data flow debug session.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Delete.json
 */
async function dataFlowDebugSessionDelete() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const request = {
    sessionId: "91fb57e0-8292-47be-89ff-c8f2d2bb2a7e",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.dataFlowDebugSession.delete(resourceGroupName, factoryName, request);
  console.log(result);
}

dataFlowDebugSessionDelete().catch(console.error);
