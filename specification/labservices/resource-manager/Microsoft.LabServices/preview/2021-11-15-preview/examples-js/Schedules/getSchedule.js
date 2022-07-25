const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the properties of a lab Schedule.
 *
 * @summary Returns the properties of a lab Schedule.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/getSchedule.json
 */
async function getSchedule() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const scheduleName = "schedule1";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.schedules.get(resourceGroupName, labName, scheduleName);
  console.log(result);
}

getSchedule().catch(console.error);
