const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Execute a data flow debug command.
 *
 * @summary Execute a data flow debug command.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_ExecuteCommand.json
 */
async function dataFlowDebugSessionExecuteCommand() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const request = {
    command: "executePreviewQuery",
    commandPayload: { rowLimits: 100, streamName: "source1" },
    sessionId: "f06ed247-9d07-49b2-b05e-2cb4a2fc871e",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.dataFlowDebugSession.beginExecuteCommandAndWait(
    resourceGroupName,
    factoryName,
    request,
  );
  console.log(result);
}
