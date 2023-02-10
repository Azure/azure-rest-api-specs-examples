const { QumuloStorage } = require("@azure/arm-liftrqumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FileSystemResource
 *
 * @summary Create a FileSystemResource
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
 */
async function fileSystemsCreateOrUpdateMinimumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "aaaaaaaaaaaaaaaaaaaaaaaa";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgopenapi";
  const fileSystemName = "aaaaaaaa";
  const resource = {
    adminPassword: "ekceujoecaashtjlsgcymnrdozk",
    delegatedSubnetId: "aaaaaaaaaa",
    initialCapacity: 9,
    location: "aaaaaaaaaaaaaaaaaaaaaaaaa",
    marketplaceDetails: {
      marketplaceSubscriptionId: "aaaaaaaaaaaaa",
      marketplaceSubscriptionStatus: "PendingFulfillmentStart",
      offerId: "aaaaaaaaaaaaaaaaaaaaaaaaa",
      planId: "aaaaaa",
      publisherId: "aa",
    },
    provisioningState: "Accepted",
    storageSku: "Standard",
    userDetails: { email: "aaaaaaaaaaaaaaaaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const result = await client.fileSystems.beginCreateOrUpdateAndWait(
    resourceGroupName,
    fileSystemName,
    resource
  );
  console.log(result);
}
