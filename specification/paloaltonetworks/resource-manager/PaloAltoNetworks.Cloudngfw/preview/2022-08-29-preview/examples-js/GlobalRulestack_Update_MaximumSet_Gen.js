const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a GlobalRulestackResource
 *
 * @summary Update a GlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/GlobalRulestack_Update_MaximumSet_Gen.json
 */
async function globalRulestackUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const globalRulestackName = "praval";
  const properties = {
    identity: {
      type: "None",
      userAssignedIdentities: {
        key16: { clientId: "aaaa", principalId: "aaaaaaaaaaaaaaa" },
      },
    },
    location: "eastus",
    properties: {
      description: "global rulestacks",
      associatedSubscriptions: ["2bf4a339-294d-4c25-b0b2-ef649e9f5c27"],
      defaultMode: "IPS",
      minAppIdVersion: "8.5.3",
      panEtag: "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
      panLocation: "eastus",
      scope: "GLOBAL",
      securityServices: {
        antiSpywareProfile: "default",
        antiVirusProfile: "default",
        dnsSubscription: "default",
        fileBlockingProfile: "default",
        outboundTrustCertificate: "default",
        outboundUnTrustCertificate: "default",
        urlFilteringProfile: "default",
        vulnerabilityProfile: "default",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.globalRulestack.update(globalRulestackName, properties);
  console.log(result);
}
