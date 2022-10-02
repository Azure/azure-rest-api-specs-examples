const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all of the private link service ids that can be linked to a Private Endpoint with auto approved in this subscription in this region.
 *
 * @summary Returns all of the private link service ids that can be linked to a Private Endpoint with auto approved in this subscription in this region.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/AutoApprovedPrivateLinkServicesResourceGroupGet.json
 */
async function getListOfPrivateLinkServiceIdThatCanBeLinkedToAPrivateEndPointWithAutoApproved() {
  const subscriptionId = "subId";
  const location = "regionName";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkServices.listAutoApprovedPrivateLinkServicesByResourceGroup(
    location,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfPrivateLinkServiceIdThatCanBeLinkedToAPrivateEndPointWithAutoApproved().catch(
  console.error
);
