const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the savings plan utilization summaries for daily or monthly grain.
 *
 * @summary Lists the savings plan utilization summaries for daily or monthly grain.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitUtilizationSummaries/SavingsPlan-SavingsPlanId-Monthly.json
 */
async function savingsPlanUtilizationSummariesMonthlyWithSavingsPlanId() {
  const savingsPlanOrderId = "66cccc66-6ccc-6c66-666c-66cc6c6c66c6";
  const savingsPlanId = "222d22dd-d2d2-2dd2-222d-2dd2222ddddd";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.benefitUtilizationSummaries.listBySavingsPlanId(
    savingsPlanOrderId,
    savingsPlanId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
