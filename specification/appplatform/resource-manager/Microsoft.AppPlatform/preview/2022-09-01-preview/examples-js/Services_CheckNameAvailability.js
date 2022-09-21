const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the resource name is valid and is not already in use.
 *
 * @summary Checks that the resource name is valid and is not already in use.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Services_CheckNameAvailability.json
 */
async function servicesCheckNameAvailability() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const availabilityParameters = {
    name: "myservice",
    type: "Microsoft.AppPlatform/Spring",
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.services.checkNameAvailability(location, availabilityParameters);
  console.log(result);
}

servicesCheckNameAvailability().catch(console.error);
