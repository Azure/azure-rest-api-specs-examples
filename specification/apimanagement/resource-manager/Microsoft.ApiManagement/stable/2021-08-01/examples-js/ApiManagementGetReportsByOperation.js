const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists report records by API Operations.
 *
 * @summary Lists report records by API Operations.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetReportsByOperation.json
 */
async function apiManagementGetReportsByOperation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const filter =
    "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reports.listByOperation(resourceGroupName, serviceName, filter)) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementGetReportsByOperation().catch(console.error);
