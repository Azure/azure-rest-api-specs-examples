const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all registered topic types.
 *
 * @summary List all registered topic types.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/TopicTypes_List.json
 */
async function topicTypesList() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.topicTypes.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
