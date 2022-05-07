Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineRunCommand() {
  const subscriptionId = "24fb23e3-6ba3-41f0-9b6e-e41131d5d61e";
  const resourceGroupName = "crptestar98131";
  const vmName = "vm3036";
  const parameters = { commandId: "RunPowerShellScript" };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginRunCommandAndWait(
    resourceGroupName,
    vmName,
    parameters
  );
  console.log(result);
}

virtualMachineRunCommand().catch(console.error);
```
