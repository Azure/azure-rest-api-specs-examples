const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Check if a resource name is available.
 *
 * @summary Check if a resource name is available.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/CheckNameAvailability.json
 */
async function checkNameAvailability() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const location = "eastus";
  const name = "accName";
  const typeParam = "Microsoft.NetApp/netAppAccounts";
  const resourceGroup = "myRG";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.netAppResource.checkNameAvailability(
    location,
    name,
    typeParam,
    resourceGroup,
  );
  console.log(result);
}
