const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists savings plan utilization summaries for the enterprise agreement scope. Supported at grain values: 'Daily' and 'Monthly'.
 *
 * @summary Lists savings plan utilization summaries for the enterprise agreement scope. Supported at grain values: 'Daily' and 'Monthly'.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitUtilizationSummaries/SavingsPlan-BillingAccount.json
 */
async function savingsPlanUtilizationSummariesBillingAccount() {
  const billingAccountId = "12345";
  const filter = "properties/usageDate ge 2022-10-15 and properties/usageDate le 2022-10-18";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.benefitUtilizationSummaries.listByBillingAccountId(
    billingAccountId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
