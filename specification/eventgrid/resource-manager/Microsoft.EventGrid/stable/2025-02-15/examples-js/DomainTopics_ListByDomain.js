const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all the topics in a domain.
 *
 * @summary List all the topics in a domain.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/DomainTopics_ListByDomain.json
 */
async function domainTopicsListByDomain() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const domainName = "exampledomain2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.domainTopics.listByDomain(resourceGroupName, domainName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
