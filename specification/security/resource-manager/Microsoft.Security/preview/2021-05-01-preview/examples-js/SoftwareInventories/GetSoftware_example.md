```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a single software data of the virtual machine.
 *
 * @summary Gets a single software data of the virtual machine.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/SoftwareInventories/GetSoftware_example.json
 */
async function getsASingleSoftwareDataOfTheVirtualMachine() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const resourceGroupName = "EITAN-TESTS";
  const resourceNamespace = "Microsoft.Compute";
  const resourceType = "virtualMachines";
  const resourceName = "Eitan-Test1";
  const softwareName = "outlook_16.0.10371.20060";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.softwareInventories.get(
    resourceGroupName,
    resourceNamespace,
    resourceType,
    resourceName,
    softwareName
  );
  console.log(result);
}

getsASingleSoftwareDataOfTheVirtualMachine().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
