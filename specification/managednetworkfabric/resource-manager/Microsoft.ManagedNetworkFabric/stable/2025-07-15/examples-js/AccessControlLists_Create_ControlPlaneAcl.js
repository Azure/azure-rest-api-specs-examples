const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to implements Access Control List PUT method.
 *
 * @summary implements Access Control List PUT method.
 * x-ms-original-file: 2025-07-15/AccessControlLists_Create_ControlPlaneAcl.json
 */
async function accessControlListsCreateMaximumSetGenGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.accessControlLists.create("example-resource-group", "example-acl", {
    annotation: "annotation",
    configurationType: "Inline",
    aclsUrl: "https://microsoft.com/a",
    defaultAction: "Permit",
    aclType: "ControlPlaneTrafficPolicy",
    deviceRole: "CE",
    controlPlaneAclConfiguration: [
      {
        ipAddressType: "IPv4",
        matchConfigurations: [
          {
            matchConfigurationName: "example-match-config",
            sequenceNumber: 3779271459,
            matchCondition: {
              protocolTypes: "tcp",
              ipCondition: { sourceIpPrefix: "10.0.0.0/24", destinationIpPrefix: "10.0.0.0/24" },
              ttlMatchCondition: { ttlValue: "1", ttlMatchType: "eq" },
              portCondition: {
                sourcePorts: { ports: ["100"], portMatchType: "eq" },
                destinationPorts: { ports: ["200"], portMatchType: "eq" },
              },
              flags: ["established"],
              icmpConfiguration: { icmpTypes: ["icmp"] },
            },
            action: { type: "Drop", remarkComment: "remark" },
          },
        ],
      },
    ],
    tags: { key5032: "example-tag-value" },
    location: "eastus",
  });
  console.log(result);
}
