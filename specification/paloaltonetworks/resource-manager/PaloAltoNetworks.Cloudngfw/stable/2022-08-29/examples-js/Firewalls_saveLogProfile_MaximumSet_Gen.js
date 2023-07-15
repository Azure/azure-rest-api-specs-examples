const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Log Profile for Firewall
 *
 * @summary Log Profile for Firewall
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/Firewalls_saveLogProfile_MaximumSet_Gen.json
 */
async function firewallsSaveLogProfileMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "firewall-rg";
  const firewallName = "firewall1";
  const logSettings = {
    applicationInsights: { id: "aaaaaaaaaaaaaaaa", key: "aaaaaaaaaaaaa" },
    commonDestination: {
      eventHubConfigurations: {
        name: "aaaaaaaa",
        id: "aaaaaaaaaa",
        nameSpace: "aaaaaaaaaaaaaaaaaaaaa",
        policyName: "aaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaaa",
      },
      monitorConfigurations: {
        id: "aaaaaaaaaaaaaaaaaaa",
        primaryKey: "aaaaaaaaaaaaa",
        secondaryKey: "a",
        subscriptionId: "aaaaaaaaaaaaa",
        workspace: "aaaaaaaaaaa",
      },
      storageConfigurations: {
        accountName: "aaaaaaaaaaaaaaaaaaaaaaa",
        id: "aaaaaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaa",
      },
    },
    decryptLogDestination: {
      eventHubConfigurations: {
        name: "aaaaaaaa",
        id: "aaaaaaaaaa",
        nameSpace: "aaaaaaaaaaaaaaaaaaaaa",
        policyName: "aaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaaa",
      },
      monitorConfigurations: {
        id: "aaaaaaaaaaaaaaaaaaa",
        primaryKey: "aaaaaaaaaaaaa",
        secondaryKey: "a",
        subscriptionId: "aaaaaaaaaaaaa",
        workspace: "aaaaaaaaaaa",
      },
      storageConfigurations: {
        accountName: "aaaaaaaaaaaaaaaaaaaaaaa",
        id: "aaaaaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaa",
      },
    },
    logOption: "SAME_DESTINATION",
    logType: "TRAFFIC",
    threatLogDestination: {
      eventHubConfigurations: {
        name: "aaaaaaaa",
        id: "aaaaaaaaaa",
        nameSpace: "aaaaaaaaaaaaaaaaaaaaa",
        policyName: "aaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaaa",
      },
      monitorConfigurations: {
        id: "aaaaaaaaaaaaaaaaaaa",
        primaryKey: "aaaaaaaaaaaaa",
        secondaryKey: "a",
        subscriptionId: "aaaaaaaaaaaaa",
        workspace: "aaaaaaaaaaa",
      },
      storageConfigurations: {
        accountName: "aaaaaaaaaaaaaaaaaaaaaaa",
        id: "aaaaaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaa",
      },
    },
    trafficLogDestination: {
      eventHubConfigurations: {
        name: "aaaaaaaa",
        id: "aaaaaaaaaa",
        nameSpace: "aaaaaaaaaaaaaaaaaaaaa",
        policyName: "aaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaaa",
      },
      monitorConfigurations: {
        id: "aaaaaaaaaaaaaaaaaaa",
        primaryKey: "aaaaaaaaaaaaa",
        secondaryKey: "a",
        subscriptionId: "aaaaaaaaaaaaa",
        workspace: "aaaaaaaaaaa",
      },
      storageConfigurations: {
        accountName: "aaaaaaaaaaaaaaaaaaaaaaa",
        id: "aaaaaaaaaaaaaaa",
        subscriptionId: "aaaaaaaaa",
      },
    },
  };
  const options = { logSettings };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.firewalls.saveLogProfile(resourceGroupName, firewallName, options);
  console.log(result);
}
