const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the collection of Certificate Authorities for the specified Gateway entity.
 *
 * @summary Lists the collection of Certificate Authorities for the specified Gateway entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGatewayCertificateAuthorities.json
 */
async function apiManagementListGatewaycertificateAuthorities() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gatewayCertificateAuthority.listByService(
    resourceGroupName,
    serviceName,
    gatewayId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
