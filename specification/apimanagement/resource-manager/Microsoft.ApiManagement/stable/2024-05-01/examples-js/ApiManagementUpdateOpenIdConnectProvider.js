const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the specific OpenID Connect Provider.
 *
 * @summary Updates the specific OpenID Connect Provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateOpenIdConnectProvider.json
 */
async function apiManagementUpdateOpenIdConnectProvider() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const opid = "templateOpenIdConnect2";
  const ifMatch = "*";
  const parameters = {
    clientSecret: "updatedsecret",
    useInApiDocumentation: true,
    useInTestConsole: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.openIdConnectProvider.update(
    resourceGroupName,
    serviceName,
    opid,
    ifMatch,
    parameters,
  );
  console.log(result);
}
