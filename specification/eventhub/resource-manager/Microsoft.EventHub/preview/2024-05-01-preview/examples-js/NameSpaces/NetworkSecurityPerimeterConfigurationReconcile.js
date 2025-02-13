const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Refreshes any information about the association.
 *
 * @summary Refreshes any information about the association.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/NetworkSecurityPerimeterConfigurationReconcile.json
 */
async function networkSecurityPerimeterConfigurationList() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "subID";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "SDK-EventHub-4794";
  const namespaceName = "sdk-Namespace-5828";
  const resourceAssociationName = "resourceAssociation1";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityPerimeterConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    namespaceName,
    resourceAssociationName,
  );
  console.log(result);
}
