const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates policy configuration for the API.
 *
 * @summary Creates or updates policy configuration for the API.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiPolicyNonXmlEncoded.json
 */
async function apiManagementCreateApiPolicyNonXmlEncoded() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "5600b57e7e8880006a040001";
  const policyId = "policy";
  const ifMatch = "*";
  const parameters = {
    format: "rawxml",
    value:
      '<policies>\r\n     <inbound>\r\n     <base />\r\n  <set-header name="newvalue" exists-action="override">\r\n   <value>"@(context.Request.Headers.FirstOrDefault(h => h.Ke=="Via"))" </value>\r\n    </set-header>\r\n  </inbound>\r\n      </policies>',
  };
  const options = { ifMatch };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiPolicy.createOrUpdate(
    resourceGroupName,
    serviceName,
    apiId,
    policyId,
    parameters,
    options
  );
  console.log(result);
}
