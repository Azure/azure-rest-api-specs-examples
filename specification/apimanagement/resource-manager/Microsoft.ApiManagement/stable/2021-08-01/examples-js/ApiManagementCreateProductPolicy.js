const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates policy configuration for the Product.
 *
 * @summary Creates or updates policy configuration for the Product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateProductPolicy.json
 */
async function apiManagementCreateProductPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "5702e97e5157a50f48dce801";
  const policyId = "policy";
  const parameters = {
    format: "xml",
    value:
      '<policies>  <inbound>    <rate-limit calls="{{call-count}}" renewal-period="15"></rate-limit>    <log-to-eventhub logger-id="16">                      @( string.Join(",", DateTime.UtcNow, context.Deployment.ServiceName, context.RequestId, context.Request.IpAddress, context.Operation.Name) )                   </log-to-eventhub>    <quota-by-key calls="40" counter-key="cc" renewal-period="3600" increment-count="@(context.Request.Method == &quot;POST&quot; ? 1:2)" />    <base />  </inbound>  <backend>    <base />  </backend>  <outbound>    <base />  </outbound></policies>',
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.productPolicy.createOrUpdate(
    resourceGroupName,
    serviceName,
    productId,
    policyId,
    parameters
  );
  console.log(result);
}

apiManagementCreateProductPolicy().catch(console.error);
