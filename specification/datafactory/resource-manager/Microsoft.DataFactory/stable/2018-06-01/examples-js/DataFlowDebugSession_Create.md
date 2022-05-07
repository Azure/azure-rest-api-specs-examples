Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.4.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function dataFlowDebugSessionCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
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
    request
  );
  console.log(result);
}

dataFlowDebugSessionCreate().catch(console.error);
```
