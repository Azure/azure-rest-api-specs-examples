const { AzureRedHatOpenShiftClient } = require("@azure/arm-redhatopenshift");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation returns properties of a SyncIdentityProvider.
 *
 * @summary The operation returns properties of a SyncIdentityProvider.
 * x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/SyncIdentityProviders_Get.json
 */
async function getsASyncIdentityProviderWithTheSpecifiedSubscriptionResourceGroupAndResourceName() {
  const subscriptionId = process.env["REDHATOPENSHIFT_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["REDHATOPENSHIFT_RESOURCE_GROUP"] || "resourceGroup";
  const resourceName = "resourceName";
  const childResourceName = "childResourceName";
  const credential = new DefaultAzureCredential();
  const client = new AzureRedHatOpenShiftClient(credential, subscriptionId);
  const result = await client.syncIdentityProviders.get(
    resourceGroupName,
    resourceName,
    childResourceName,
  );
  console.log(result);
}
