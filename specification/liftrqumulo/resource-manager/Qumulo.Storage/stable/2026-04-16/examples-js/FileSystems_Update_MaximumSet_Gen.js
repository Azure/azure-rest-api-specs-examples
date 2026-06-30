const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a FileSystemResource
 *
 * @summary update a FileSystemResource
 * x-ms-original-file: 2026-04-16/FileSystems_Update_MaximumSet_Gen.json
 */
async function fileSystemsUpdateMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "C9CC2D2A-5AA0-4839-A85F-18491F2D244A";
  const client = new QumuloStorage(credential, subscriptionId);
  const result = await client.fileSystems.update("rgQumulo", "qumulo-fs-01", {
    identity: { type: "None", userAssignedIdentities: { key8111: {} } },
    tags: { key5846: "pyrehicqychbuvhcqcbbstzcxsueyf" },
    properties: {
      marketplaceDetails: {
        marketplaceSubscriptionId: "vwjzkiurjihwxrhoicenkbxacokvep",
        planId: "vxnyxa",
        offerId: "itiocfnteqyuavgmdtnvwvbpectyr",
        publisherId: "zfevjvhjiifwxbazta",
        termUnit: "lkbiqoqdyqbua",
      },
      userDetails: { email: "rlqqzevfgtqpynvifqp" },
      delegatedSubnetId: "osinzkhurmzdaw",
      performanceTier: "Premium",
    },
  });
  console.log(result);
}
