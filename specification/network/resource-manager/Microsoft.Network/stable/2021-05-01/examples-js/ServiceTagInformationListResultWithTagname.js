const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of service tag information resources with pagination.
 *
 * @summary Gets a list of service tag information resources with pagination.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ServiceTagInformationListResultWithTagname.json
 */
async function getListOfServiceTagsWithTagName() {
  const subscriptionId = "subid";
  const location = "westeurope";
  const tagName = "ApiManagement";
  const options = { tagName };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceTagInformationOperations.list(location, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfServiceTagsWithTagName().catch(console.error);
