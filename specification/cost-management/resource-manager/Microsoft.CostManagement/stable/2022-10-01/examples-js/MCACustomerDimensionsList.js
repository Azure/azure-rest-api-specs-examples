const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the dimensions by the defined scope.
 *
 * @summary Lists the dimensions by the defined scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/MCACustomerDimensionsList.json
 */
async function customerDimensionsListMca() {
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/5678";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.dimensions.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
