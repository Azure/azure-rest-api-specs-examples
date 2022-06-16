const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of service tag information resources with pagination.
 *
 * @summary Gets a list of service tag information resources with pagination.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ServiceTagInformationListResult.json
 */
async function getListOfServiceTags() {
  const subscriptionId = "subid";
  const location = "westeurope";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceTagInformationOperations.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfServiceTags().catch(console.error);
