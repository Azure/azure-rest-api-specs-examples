const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Synchronously updates a channel with the specified parameters.
 *
 * @summary Synchronously updates a channel with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Channels_Update.json
 */
async function channelsUpdate() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
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
