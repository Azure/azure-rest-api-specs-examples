const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the give namespace name availability.
 *
 * @summary Check the give namespace name availability.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasCheckNameAvailability.json
 */
async function aliasNameAvailability() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-9080";
  const parameters = {
    name: "sdk-DisasterRecovery-9474",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.checkNameAvailability(
    resourceGroupName,
    namespaceName,
    parameters
  );
  console.log(result);
}
