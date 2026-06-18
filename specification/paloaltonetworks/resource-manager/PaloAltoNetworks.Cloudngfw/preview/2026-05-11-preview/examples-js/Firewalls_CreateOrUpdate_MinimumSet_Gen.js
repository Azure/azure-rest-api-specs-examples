const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a FirewallResource
 *
 * @summary create a FirewallResource
 * x-ms-original-file: 2026-05-11-preview/Firewalls_CreateOrUpdate_MinimumSet_Gen.json
 */
async function firewallsCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.firewalls.createOrUpdate("firewall-rg", "firewall1", {
    location: "eastus",
    dnsSettings: {},
    marketplaceDetails: { offerId: "liftr-pan-ame-test", publisherId: "isvtestuklegacy" },
    networkProfile: {
      enableEgressNat: "ENABLED",
      networkType: "VNET",
      publicIps: [
        {
          address: "20.22.92.11",
          resourceId:
            "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1",
        },
      ],
    },
    planData: { billingCycle: "MONTHLY", planId: "liftrpantestplan" },
  });
  console.log(result);
}
