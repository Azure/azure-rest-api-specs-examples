const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a GlobalRulestackResource
 *
 * @summary Create a GlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/GlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function globalRulestackCreateOrUpdateMaximumSetGen() {
  const globalRulestackName = "praval";
  const resource = {
    description: "global rulestacks",
    associatedSubscriptions: ["2bf4a339-294d-4c25-b0b2-ef649e9f5c27"],
    defaultMode: "IPS",
    identity: {
      type: "None",
      userAssignedIdentities: {
        key16: { clientId: "aaaa", principalId: "aaaaaaaaaaaaaaa" },
      },
    },
    location: "eastus",
    minAppIdVersion: "8.5.3",
    panEtag: "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
    panLocation: "eastus",
    provisioningState: "Accepted",
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
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    resource
  );
  console.log(result);
}
