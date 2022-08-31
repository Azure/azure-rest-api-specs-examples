const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a Lab Plan resource. Deleting a lab plan does not delete labs associated with a lab plan, nor does it delete shared images added to a gallery via the lab plan permission container.
 *
 * @summary Operation to delete a Lab Plan resource. Deleting a lab plan does not delete labs associated with a lab plan, nor does it delete shared images added to a gallery via the lab plan permission container.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/deleteLabPlan.json
 */
async function deleteLabPlan() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labPlanName = "testlabplan";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.labPlans.beginDeleteAndWait(resourceGroupName, labPlanName);
  console.log(result);
}

deleteLabPlan().catch(console.error);
