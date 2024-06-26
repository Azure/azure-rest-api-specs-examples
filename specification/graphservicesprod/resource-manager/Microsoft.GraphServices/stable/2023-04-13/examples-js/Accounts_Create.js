const { GraphServices } = require("@azure/arm-graphservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update account resource.
 *
 * @summary Create or update account resource.
 * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_Create.json
 */
async function createAccountResource() {
  const subscriptionId =
    process.env["GRAPHSERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["GRAPHSERVICES_RESOURCE_GROUP"] || "testResourceGroupGRAM";
  const resourceName = "11111111-aaaa-1111-bbbb-1111111111111";
  const accountResource = {
    properties: { appId: "11111111-aaaa-1111-bbbb-111111111111" },
  };
  const credential = new DefaultAzureCredential();
  const client = new GraphServices(credential, subscriptionId);
  const result = await client.accounts.beginCreateAndUpdateAndWait(
    resourceGroupName,
    resourceName,
    accountResource
  );
  console.log(result);
}
