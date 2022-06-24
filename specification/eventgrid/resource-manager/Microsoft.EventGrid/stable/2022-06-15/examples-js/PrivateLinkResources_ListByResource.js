const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the private link resources under a topic, domain, or partner namespace.
 *
 * @summary List all the private link resources under a topic, domain, or partner namespace.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PrivateLinkResources_ListByResource.json
 */
async function privateLinkResourcesListByResource() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const parentType = "topics";
  const parentName = "exampletopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listByResource(
    resourceGroupName,
    parentType,
    parentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateLinkResourcesListByResource().catch(console.error);
