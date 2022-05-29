Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearning_1.0.0-beta.1/sdk/machinelearning/arm-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns supported VM Sizes in a location
 *
 * @summary Returns supported VM Sizes in a location
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/VirtualMachineSize/list.json
 */
async function listVMSizes() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.virtualMachineSizes.list(location);
  console.log(result);
}

listVMSizes().catch(console.error);
```
