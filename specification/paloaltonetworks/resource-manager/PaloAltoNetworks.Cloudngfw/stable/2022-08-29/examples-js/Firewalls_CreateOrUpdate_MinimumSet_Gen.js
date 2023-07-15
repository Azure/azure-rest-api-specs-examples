const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FirewallResource
 *
 * @summary Create a FirewallResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/Firewalls_CreateOrUpdate_MinimumSet_Gen.json
 */
async function firewallsCreateOrUpdateMinimumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "firewall-rg";
  const firewallName = "firewall1";
  const resource = {
    dnsSettings: {},
    location: "eastus",
    marketplaceDetails: {
      offerId: "liftr-pan-ame-test",
      publisherId: "isvtestuklegacy",
    },
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
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.firewalls.beginCreateOrUpdateAndWait(
    resourceGroupName,
    firewallName,
    resource
  );
  console.log(result);
}
