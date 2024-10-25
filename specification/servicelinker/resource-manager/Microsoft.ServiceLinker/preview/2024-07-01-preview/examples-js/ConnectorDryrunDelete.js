const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a dryrun job
 *
 * @summary delete a dryrun job
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/ConnectorDryrunDelete.json
 */
async function connectorDryrunDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICELINKER_RESOURCE_GROUP"] || "test-rg";
  const location = "westus";
  const dryrunName = "dryrunName";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.connector.deleteDryrun(
    subscriptionId,
    resourceGroupName,
    location,
    dryrunName,
  );
  console.log(result);
}
