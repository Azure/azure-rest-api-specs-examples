Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.4.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function factoriesGetDataPlaneAccess() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const policy = {
    accessResourcePath: "",
    expireTime: "2018-11-10T09:46:20.2659347Z",
    permissions: "r",
    profileName: "DefaultProfile",
    startTime: "2018-11-10T02:46:20.2659347Z",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.factories.getDataPlaneAccess(resourceGroupName, factoryName, policy);
  console.log(result);
}

factoriesGetDataPlaneAccess().catch(console.error);
```
