const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a FirewallResource
 *
 * @summary Update a FirewallResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/Firewalls_Update_MaximumSet_Gen.json
 */
async function firewallsUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "firewall-rg";
  const firewallName = "firewall1";
  const properties = {
    identity: {
      type: "None",
      userAssignedIdentities: {
        key16: { clientId: "aaaa", principalId: "aaaaaaaaaaaaaaa" },
      },
    },
    properties: {
      associatedRulestack: {
        location: "eastus",
        resourceId: "aaaaaaaaaa",
        rulestackId: "aaaaaaaaaaaaaaaa",
      },
      dnsSettings: {
        dnsServers: [
          {
            address: "20.22.92.111",
            resourceId:
              "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1",
          },
        ],
        enableDnsProxy: "DISABLED",
        enabledDnsType: "CUSTOM",
      },
      frontEndSettings: [
        {
          name: "frontendsetting11",
          backendConfiguration: {
            address: {
              address: "20.22.32.136",
              resourceId:
                "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp2",
            },
            port: "80",
          },
          frontendConfiguration: {
            address: {
              address: "20.22.91.251",
              resourceId:
                "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp1",
            },
            port: "80",
          },
          protocol: "TCP",
        },
      ],
      isPanoramaManaged: "TRUE",
      marketplaceDetails: {
        marketplaceSubscriptionStatus: "PendingFulfillmentStart",
        offerId: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        publisherId: "aaaa",
      },
      networkProfile: {
        egressNatIp: [
          {
            address: "20.22.92.111",
            resourceId:
              "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1",
          },
        ],
        enableEgressNat: "ENABLED",
        networkType: "VNET",
        publicIps: [
          {
            address: "20.22.92.11",
            resourceId:
              "/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1",
          },
        ],
        vnetConfiguration: {
          ipOfTrustSubnetForUdr: {
            address: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
          },
          trustSubnet: {
            addressSpace: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet",
          },
          unTrustSubnet: {
            addressSpace: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
          },
          vnet: {
            addressSpace: "10.1.0.0/16",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet",
          },
        },
        vwanConfiguration: {
          ipOfTrustSubnetForUdr: {
            address: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
          },
          networkVirtualApplianceId: "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
          trustSubnet: {
            addressSpace: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet",
          },
          unTrustSubnet: {
            addressSpace: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
          },
          vHub: {
            addressSpace: "10.1.1.0/24",
            resourceId:
              "/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet",
          },
        },
      },
      panEtag: "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
      panoramaConfig: { configString: "bas64EncodedString" },
      planData: {
        billingCycle: "WEEKLY",
        planId: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        usageType: "PAYG",
      },
    },
    tags: { tagName: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.firewalls.update(resourceGroupName, firewallName, properties);
  console.log(result);
}
