const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if a quota is available.
 *
 * @summary Check if a quota is available.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2024-07-01-preview/examples/CheckQuotaAvailability.json
 */
async function checkQuotaAvailability() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const location = "eastus";
  const name = "resource1";
  const typeParam = "Microsoft.NetApp/netAppAccounts";
  const resourceGroup = "myRG";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.netAppResource.checkQuotaAvailability(
    location,
    name,
    typeParam,
    resourceGroup,
  );
  console.log(result);
}
