const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete existing domain topic.
 *
 * @summary Delete existing domain topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainTopics_Delete.json
 */
async function domainTopicsDelete() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const domainName = "exampledomain1";
  const domainTopicName = "exampledomaintopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.domainTopics.beginDeleteAndWait(
    resourceGroupName,
    domainName,
    domainTopicName
  );
  console.log(result);
}

domainTopicsDelete().catch(console.error);
