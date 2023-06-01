const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FileSystemResource
 *
 * @summary Create a FileSystemResource
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
 */
async function fileSystemsCreateOrUpdateMaximumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "ulseeqylxb";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgQumulo";
  const fileSystemName = "nauwwbfoqehgbhdsmkewoboyxeqg";
  const resource = {
    adminPassword: "ekceujoecaashtjlsgcymnrdozk",
    availabilityZone: "maseyqhlnhoiwbabcqabtedbjpip",
    clusterLoginUrl: "jjqhgevy",
    delegatedSubnetId: "neqctctqdmjezfgt",
    identity: { type: "None", userAssignedIdentities: { key4522: {} } },
    initialCapacity: 9,
    location: "przdlsmlzsszphnixq",
    marketplaceDetails: {
      marketplaceSubscriptionId: "ujrcqvxfnhxxheoth",
      marketplaceSubscriptionStatus: "PendingFulfillmentStart",
      offerId: "eiyhbmpwgezcmzrrfoiskuxlcvwojf",
      planId: "x",
      publisherId: "wfmokfdjbwpjhz",
    },
    privateIPs: ["kslguxrwbwkrj"],
    provisioningState: "Accepted",
    storageSku: "Standard",
    tags: { key6565: "cgdhmupta" },
    userDetails: { email: "viptslwulnpaupfljvnjeq" },
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
