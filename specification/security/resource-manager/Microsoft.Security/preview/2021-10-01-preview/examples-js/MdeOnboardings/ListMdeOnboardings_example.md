```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The configuration or data needed to onboard the machine to MDE
 *
 * @summary The configuration or data needed to onboard the machine to MDE
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-10-01-preview/examples/MdeOnboardings/ListMdeOnboardings_example.json
 */
async function theConfigurationOrDataNeededToOnboardTheMachineToMde() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.mdeOnboardings.list();
  console.log(result);
}

theConfigurationOrDataNeededToOnboardTheMachineToMde().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
