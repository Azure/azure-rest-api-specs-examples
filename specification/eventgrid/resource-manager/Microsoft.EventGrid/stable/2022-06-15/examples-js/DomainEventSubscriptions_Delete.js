const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an existing event subscription for a domain.
 *
 * @summary Delete an existing event subscription for a domain.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainEventSubscriptions_Delete.json
 */
async function domainEventSubscriptionsDelete() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const domainName = "exampleDomain1";
  const eventSubscriptionName = "examplesubscription1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.domainEventSubscriptions.beginDeleteAndWait(
    resourceGroupName,
    domainName,
    eventSubscriptionName
  );
  console.log(result);
}

domainEventSubscriptionsDelete().catch(console.error);
