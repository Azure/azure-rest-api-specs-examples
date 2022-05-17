Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.5.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function dataFlowDebugSessionExecuteCommand() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
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
    request
  );
  console.log(result);
}

dataFlowDebugSessionExecuteCommand().catch(console.error);
```
