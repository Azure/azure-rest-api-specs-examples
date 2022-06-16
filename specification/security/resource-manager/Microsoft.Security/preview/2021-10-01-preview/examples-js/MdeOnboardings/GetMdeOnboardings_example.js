const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The default configuration or data needed to onboard the machine to MDE
 *
 * @summary The default configuration or data needed to onboard the machine to MDE
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-10-01-preview/examples/MdeOnboardings/GetMdeOnboardings_example.json
 */
async function theDefaultConfigurationOrDataNeededToOnboardTheMachineToMde() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.mdeOnboardings.get();
  console.log(result);
}

theDefaultConfigurationOrDataNeededToOnboardTheMachineToMde().catch(console.error);
