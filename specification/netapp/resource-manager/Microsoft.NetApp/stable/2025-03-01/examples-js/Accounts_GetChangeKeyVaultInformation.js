const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Contains data from encryption.keyVaultProperties as well as information about which private endpoint is used by each encryption sibling set. Response from this endpoint can be modified and used as request body for POST request.
 *
 * @summary Contains data from encryption.keyVaultProperties as well as information about which private endpoint is used by each encryption sibling set. Response from this endpoint can be modified and used as request body for POST request.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/Accounts_GetChangeKeyVaultInformation.json
 */
async function accountsGetChangeKeyVaultInformation() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginGetChangeKeyVaultInformationAndWait(
    resourceGroupName,
    accountName,
  );
  console.log(result);
}
