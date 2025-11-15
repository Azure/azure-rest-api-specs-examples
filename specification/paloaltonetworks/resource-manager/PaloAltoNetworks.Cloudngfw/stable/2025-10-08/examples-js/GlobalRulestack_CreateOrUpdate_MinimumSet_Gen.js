const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a GlobalRulestackResource
 *
 * @summary Create a GlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2025-10-08/examples/GlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function globalRulestackCreateOrUpdateMinimumSetGen() {
  const globalRulestackName = "praval";
  const resource = { location: "eastus" };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    resource,
  );
  console.log(result);
}
