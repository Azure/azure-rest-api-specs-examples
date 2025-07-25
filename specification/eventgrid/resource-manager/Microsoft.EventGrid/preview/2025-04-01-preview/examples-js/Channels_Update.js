const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Synchronously updates a channel with the specified parameters.
 *
 * @summary Synchronously updates a channel with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Channels_Update.json
 */
async function channelsUpdate() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const partnerNamespaceName = "examplePartnerNamespaceName1";
  const channelName = "exampleChannelName1";
  const channelUpdateParameters = {
    expirationTimeIfNotActivatedUtc: new Date("2022-03-23T23:06:11.785Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.channels.update(
    resourceGroupName,
    partnerNamespaceName,
    channelName,
    channelUpdateParameters,
  );
  console.log(result);
}
