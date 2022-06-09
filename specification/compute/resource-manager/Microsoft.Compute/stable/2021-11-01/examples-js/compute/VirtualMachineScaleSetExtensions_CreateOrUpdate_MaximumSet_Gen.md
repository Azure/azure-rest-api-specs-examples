```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetExtensionsCreateOrUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaa";
  const vmssExtensionName = "aaaaaaaaaaaaaaaaaaaaa";
  const extensionParameters = {
    name: "{extension-name}",
    typePropertiesType: "{extension-Type}",
    autoUpgradeMinorVersion: true,
    enableAutomaticUpgrade: true,
    forceUpdateTag: "aaaaaaaaa",
    protectedSettings: {},
    provisionAfterExtensions: ["aa"],
    publisher: "{extension-Publisher}",
    settings: {},
    suppressFailures: true,
    typeHandlerVersion: "{handler-version}",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName,
    extensionParameters
  );
  console.log(result);
}

virtualMachineScaleSetExtensionsCreateOrUpdateMaximumSetGen().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.
