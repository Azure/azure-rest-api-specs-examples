const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a GlobalRulestackResource
 *
 * @summary create a GlobalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function globalRulestackCreateOrUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.createOrUpdate("praval", {
    identity: {
      type: "None",
      userAssignedIdentities: { key16: { clientId: "aaaa", principalId: "aaaaaaaaaaaaaaa" } },
    },
    location: "eastus",
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
  });
  console.log(result);
}
