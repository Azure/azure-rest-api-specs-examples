const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a data flow debug session.
 *
 * @summary Creates a data flow debug session.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Create.json
 */
async function dataFlowDebugSessionCreate() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const request = {
    integrationRuntime: {
      name: "ir1",
      properties: {
        type: "Managed",
        computeProperties: {
          dataFlowProperties: {
            computeType: "General",
            coreCount: 48,
            timeToLive: 10,
          },
          location: "AutoResolve",
        },
      },
    },
    timeToLive: 60,
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.dataFlowDebugSession.beginCreateAndWait(
    resourceGroupName,
    factoryName,
    request,
  );
  console.log(result);
}
