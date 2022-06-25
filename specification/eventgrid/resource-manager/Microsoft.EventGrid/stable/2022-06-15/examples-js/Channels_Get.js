const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get properties of a channel.
 *
 * @summary Get properties of a channel.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Channels_Get.json
 */
async function channelsGet() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const partnerNamespaceName = "examplePartnerNamespaceName1";
  const channelName = "exampleChannelName1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.channels.get(resourceGroupName, partnerNamespaceName, channelName);
  console.log(result);
}

channelsGet().catch(console.error);
