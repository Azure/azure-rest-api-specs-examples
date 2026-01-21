const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a FileSystemResource
 *
 * @summary create a FileSystemResource
 * x-ms-original-file: 2025-03-21/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
 */
async function fileSystemsCreateOrUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4B6E265D-57CF-4A9D-8B35-3CC68ED9D208";
  const client = new StorageClient(credential, subscriptionId);
  const result = await client.fileSystems.createOrUpdate("rgDell", "abcd", {
    properties: {
      marketplace: {
        marketplaceSubscriptionId: "mvjcxwndudbylynme",
        planId: "eekvwfndjoxijeasksnt",
        offerId: "bcganbkmvznyqfnvhjuag",
        publisherId: "trdzykoeskmcwpo",
        marketplaceSubscriptionStatus: "PendingFulfillmentStart",
        privateOfferId: "privateOfferId",
        planName: "planeName",
      },
      delegatedSubnetId: "rqkpvczbtqcxiaivtbuixblb",
      delegatedSubnetCidr: "10.0.0.1/24",
      user: { email: "jwogfgznmjabdbcjcljjlkxdpc" },
      smartConnectFqdn: "fqdn",
      oneFsUrl: "oneFsUrl",
      dellReferenceNumber: "fhewkj",
      encryption: {
        encryptionType: "Customer-managed keys (CMK)",
        keyUrl: "https://contoso.com/keyurl/keyVersion",
        encryptionIdentityProperties: {
          identityType: "UserAssigned",
          identityResourceId:
            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}",
        },
      },
    },
    identity: { type: "UserAssigned", userAssignedIdentities: { key7644: {} } },
    tags: { key7594: "sfkwapubiurgedzveido" },
    location: "cvbmsqftppe",
  });
  console.log(result);
}
