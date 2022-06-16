const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a security connector.
 *
 * @summary Deletes a security connector.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/SecurityConnectors/DeleteSecurityConnector_example.json
 */
async function deleteASecurityConnector() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityConnectors.delete(resourceGroupName, securityConnectorName);
  console.log(result);
}

deleteASecurityConnector().catch(console.error);
