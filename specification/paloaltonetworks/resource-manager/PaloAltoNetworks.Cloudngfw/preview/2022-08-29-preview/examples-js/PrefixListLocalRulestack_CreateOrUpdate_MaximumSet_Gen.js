const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a PrefixListResource
 *
 * @summary Create a PrefixListResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PrefixListLocalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function prefixListLocalRulestackCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "rgopenapi";
  const localRulestackName = "lrs1";
  const name = "armid1";
  const resource = {
    description: "string",
    auditComment: "comment",
    etag: "2bf4a339-294d-4c25-b0b2-ef649e9f5c27",
    prefixList: ["1.0.0.0/24"],
    provisioningState: "Accepted",
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.prefixListLocalRulestack.beginCreateOrUpdateAndWait(
    resourceGroupName,
    localRulestackName,
    name,
    resource
  );
  console.log(result);
}
