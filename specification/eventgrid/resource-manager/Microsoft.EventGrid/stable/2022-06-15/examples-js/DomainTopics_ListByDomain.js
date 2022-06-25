const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the topics in a domain.
 *
 * @summary List all the topics in a domain.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainTopics_ListByDomain.json
 */
async function domainTopicsListByDomain() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const domainName = "exampledomain2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.domainTopics.listByDomain(resourceGroupName, domainName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

domainTopicsListByDomain().catch(console.error);
