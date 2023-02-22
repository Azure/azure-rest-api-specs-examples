const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all migrationConfigurations
 *
 * @summary Gets all migrationConfigurations
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Migrationconfigurations/SBMigrationconfigurationList.json
 */
async function migrationConfigurationsList() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "SubscriptionId";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-9259";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.migrationConfigs.list(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
