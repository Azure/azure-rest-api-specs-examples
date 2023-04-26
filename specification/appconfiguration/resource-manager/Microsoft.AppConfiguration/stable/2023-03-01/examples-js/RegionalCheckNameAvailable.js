const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the configuration store name is available for use.
 *
 * @summary Checks whether the configuration store name is available for use.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/RegionalCheckNameAvailable.json
 */
async function configurationStoresCheckNameAvailable() {
  const subscriptionId =
    process.env["APPCONFIGURATION_SUBSCRIPTION_ID"] || "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const location = "westus";
  const checkNameAvailabilityParameters = {
    name: "contoso",
    type: "Microsoft.AppConfiguration/configurationStores",
  };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.operations.regionalCheckNameAvailability(
    location,
    checkNameAvailabilityParameters
  );
  console.log(result);
}
