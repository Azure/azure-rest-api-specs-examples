const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the specified NetApp account within the resource group
 *
 * @summary create or update the specified NetApp account within the resource group
 * x-ms-original-file: 2025-12-15-preview/Accounts_CreateOrUpdateLdapConfig.json
 */
async function accountsCreateOrUpdateLdapConfig() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.accounts.createOrUpdate("myRG", "account1", {
    location: "eastus",
    properties: {
      ldapConfiguration: {
        domain: "example.com",
        ldapServers: ["192.0.2.1", "192.0.2.2"],
        ldapOverTLS: false,
      },
    },
  });
  console.log(result);
}
