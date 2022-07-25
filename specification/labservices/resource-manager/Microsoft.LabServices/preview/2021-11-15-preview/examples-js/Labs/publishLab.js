const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Publish or re-publish a lab. This will create or update all lab resources, such as virtual machines.
 *
 * @summary Publish or re-publish a lab. This will create or update all lab resources, such as virtual machines.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/publishLab.json
 */
async function publishLab() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.labs.beginPublishAndWait(resourceGroupName, labName);
  console.log(result);
}

publishLab().catch(console.error);
