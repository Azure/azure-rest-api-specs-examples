const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of all possible traffic between resources for the subscription and location, based on connection type.
 *
 * @summary Gets the list of all possible traffic between resources for the subscription and location, based on connection type.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AllowedConnections/GetAllowedConnections_example.json
 */
async function getAllowedConnections() {
  const subscriptionId = "3eeab341-f466-499c-a8be-85427e154bad";
  const resourceGroupName = "myResourceGroup";
  const ascLocation = "centralus";
  const connectionType = "Internal";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.allowedConnections.get(
    resourceGroupName,
    ascLocation,
    connectionType
  );
  console.log(result);
}

getAllowedConnections().catch(console.error);
