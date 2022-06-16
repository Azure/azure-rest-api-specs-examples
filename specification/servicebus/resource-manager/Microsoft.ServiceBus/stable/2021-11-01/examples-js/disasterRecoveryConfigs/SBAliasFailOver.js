const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
 *
 * @summary Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs/SBAliasFailOver.json
 */
async function sbAliasFailOver() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ardsouzatestRG";
  const namespaceName = "sdk-Namespace-8860";
  const alias = "sdk-DisasterRecovery-3814";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.failOver(
    resourceGroupName,
    namespaceName,
    alias
  );
  console.log(result);
}

sbAliasFailOver().catch(console.error);
