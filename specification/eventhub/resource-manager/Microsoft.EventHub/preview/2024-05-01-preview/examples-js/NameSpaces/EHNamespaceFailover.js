const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to GeoDR Failover
 *
 * @summary GeoDR Failover
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/EHNamespaceFailover.json
 */
async function nameSpaceCreate() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "SampleSubscription";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ResurceGroupSample";
  const namespaceName = "NamespaceGeoDRFailoverSample";
  const parameters = { force: true, primaryLocation: "centralus" };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.beginFailoverAndWait(
    resourceGroupName,
    namespaceName,
    parameters,
  );
  console.log(result);
}
