const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a FileSystemResource
 *
 * @summary create a FileSystemResource
 * x-ms-original-file: 2025-03-21/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
 */
async function fileSystemsCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "BF7E7352-2FE4-4163-9CF7-5FF8EC2E9B92";
  const client = new StorageClient(credential, subscriptionId);
  const result = await client.fileSystems.createOrUpdate("rgDell", "abcd", {
    properties: {
      marketplace: {
        planId: "lgozf",
        offerId: "pzhjvibxqgeqkndqnjlduwnxqbr",
        privateOfferId: "privateOfferId",
        planName: "planeName",
      },
      delegatedSubnetId: "yp",
      delegatedSubnetCidr: "10.0.0.1/24",
      user: { email: "hoznewwtzmyjzctzosfuh" },
      dellReferenceNumber: "fhewkj",
      encryption: { encryptionType: "Microsoft-managed keys (MMK)" },
    },
    location: "tbcvhxzpgrijtdygkttnfswwtacs",
  });
  console.log(result);
}
