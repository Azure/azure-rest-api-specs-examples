const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of the APIs associated with a gateway.
 *
 * @summary Lists a collection of the APIs associated with a gateway.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListGatewayApis.json
 */
async function apiManagementListGatewayApis() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gatewayApi.listByService(
    resourceGroupName,
    serviceName,
    gatewayId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListGatewayApis().catch(console.error);
