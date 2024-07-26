const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FileSystemResource
 *
 * @summary Create a FileSystemResource
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
 */
async function fileSystemsCreateOrUpdateMinimumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "aaaaaaaaaaaaaaaaaaaaaaaa";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgopenapi";
  const fileSystemName = "aaaaaaaa";
  const resource = {
    adminPassword: "fakeTestSecretPlaceholder",
    delegatedSubnetId: "aaaaaaaaaa",
    location: "aaaaaaaaaaaaaaaaaaaaaaaaa",
    marketplaceDetails: {
      marketplaceSubscriptionId: "aaaaaaaaaaaaa",
      marketplaceSubscriptionStatus: "PendingFulfillmentStart",
      offerId: "aaaaaaaaaaaaaaaaaaaaaaaaa",
      planId: "aaaaaa",
    },
    storageSku: "Standard",
    userDetails: { email: "viptslwulnpaupfljvnjeq" },
  };
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const result = await client.fileSystems.beginCreateOrUpdateAndWait(
    resourceGroupName,
    fileSystemName,
    resource,
  );
  console.log(result);
}
