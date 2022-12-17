const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if the specified Visual Studio Team Services account name is available. Resource name can be either an account name or an account name and PUID.
 *
 * @summary Checks if the specified Visual Studio Team Services account name is available. Resource name can be either an account name or an account name and PUID.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/CheckNameAvailability.json
 */
async function checkAvailabilityOfAnAccountName() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const body = {
    resourceName: "ExampleName",
    resourceType: "Account",
  };
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.accounts.checkNameAvailability(body);
  console.log(result);
}

checkAvailabilityOfAnAccountName().catch(console.error);
