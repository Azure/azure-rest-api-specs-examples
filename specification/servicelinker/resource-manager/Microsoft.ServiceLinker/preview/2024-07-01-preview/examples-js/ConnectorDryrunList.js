const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list dryrun jobs
 *
 * @summary list dryrun jobs
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/ConnectorDryrunList.json
 */
async function connectorDryrunList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICELINKER_RESOURCE_GROUP"] || "test-rg";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.connector.listDryrun(subscriptionId, resourceGroupName, location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
