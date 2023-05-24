const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the result of the specified operation. The link with this operationId is provided as a response header of the initial request.
 *
 * @summary Gets the result of the specified operation. The link with this operationId is provided as a response header of the initial request.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportOperationResultsBySubscriptionScope.json
 */
async function getDetailsOfTheOperationResult() {
  const operationId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.generateDetailedCostReportOperationResults.beginGetAndWait(
    operationId,
    scope
  );
  console.log(result);
}
