const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the result of the specified operation. This link is provided in the CostDetails creation request response Location header.
 *
 * @summary Get the result of the specified operation. This link is provided in the CostDetails creation request response Location header.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/CostDetailsOperationResultsBySubscriptionScope.json
 */
async function getDetailsOfTheOperationResult() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.generateCostDetailsReport.beginGetOperationResultsAndWait(
    scope,
    operationId
  );
  console.log(result);
}
