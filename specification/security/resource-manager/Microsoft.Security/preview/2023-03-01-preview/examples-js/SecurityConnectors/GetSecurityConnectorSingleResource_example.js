const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves details of a specific security connector
 *
 * @summary Retrieves details of a specific security connector
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorSingleResource_example.json
 */
async function retrieveASecurityConnector() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "a5caac9c-5c04-49af-b3d0-e204f40345d5";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const securityConnectorName = "exampleSecurityConnectorName";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityConnectors.get(resourceGroupName, securityConnectorName);
  console.log(result);
}
