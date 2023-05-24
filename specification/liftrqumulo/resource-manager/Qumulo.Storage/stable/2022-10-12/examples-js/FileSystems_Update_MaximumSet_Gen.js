const { QumuloStorage } = require("@azure/arm-liftrqumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a FileSystemResource
 *
 * @summary Update a FileSystemResource
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Update_MaximumSet_Gen.json
 */
async function fileSystemsUpdateMaximumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "ulseeqylxb";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgQumulo";
  const fileSystemName = "nauwwbfoqehgbhdsmkewoboyxeqg";
  const properties = {
    identity: { type: "None", userAssignedIdentities: { key4522: {} } },
    properties: {
      clusterLoginUrl: "adabmuthwrbjshzfbo",
      delegatedSubnetId: "vjfirtaljehawmflyfianw",
      marketplaceDetails: {
        marketplaceSubscriptionId: "ujrcqvxfnhxxheoth",
        marketplaceSubscriptionStatus: "PendingFulfillmentStart",
        offerId: "eiyhbmpwgezcmzrrfoiskuxlcvwojf",
        planId: "x",
        publisherId: "wfmokfdjbwpjhz",
      },
      privateIPs: ["eugjqbaoucgjsopzfrq"],
      userDetails: { email: "aa" },
    },
    tags: { key7534: "jsgqvqbagquvxowbrkanyhzvo" },
  };
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const result = await client.fileSystems.update(resourceGroupName, fileSystemName, properties);
  console.log(result);
}
