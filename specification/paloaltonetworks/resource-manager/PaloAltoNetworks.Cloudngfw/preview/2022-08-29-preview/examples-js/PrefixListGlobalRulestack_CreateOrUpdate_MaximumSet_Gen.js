const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a PrefixListGlobalRulestackResource
 *
 * @summary Create a PrefixListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PrefixListGlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function prefixListGlobalRulestackCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const globalRulestackName = "praval";
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
  const result = await client.prefixListGlobalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    name,
    resource
  );
  console.log(result);
}
