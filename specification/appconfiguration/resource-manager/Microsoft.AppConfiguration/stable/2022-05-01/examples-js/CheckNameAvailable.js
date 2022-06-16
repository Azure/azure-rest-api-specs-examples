const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the configuration store name is available for use.
 *
 * @summary Checks whether the configuration store name is available for use.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/CheckNameAvailable.json
 */
async function configurationStoresCheckNameAvailable() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const checkNameAvailabilityParameters = {
    name: "contoso",
    type: "Microsoft.AppConfiguration/configurationStores",
  };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.operations.checkNameAvailability(checkNameAvailabilityParameters);
  console.log(result);
}

configurationStoresCheckNameAvailable().catch(console.error);
