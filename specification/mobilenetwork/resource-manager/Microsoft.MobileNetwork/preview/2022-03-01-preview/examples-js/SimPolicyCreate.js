const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a SimPolicy.
 *
 * @summary Creates or updates a SimPolicy.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyCreate.json
 */
async function createSimPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const simPolicyName = "testPolicy";
  const parameters = {
    defaultSlice: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice",
    },
    location: "eastus",
    registrationTimer: 3240,
    sliceConfigurations: [
      {
        dataNetworkConfigurations: [
          {
            fiveQi: 9,
            additionalAllowedSessionTypes: [],
            allocationAndRetentionPriorityLevel: 9,
            allowedServices: [
              {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/testService",
              },
            ],
            dataNetwork: {
              id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork",
            },
            defaultSessionType: "IPv4",
            preemptionCapability: "NotPreempt",
            preemptionVulnerability: "Preemptable",
            sessionAmbr: { downlink: "1 Gbps", uplink: "500 Mbps" },
          },
        ],
        defaultDataNetwork: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork",
        },
        slice: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice",
        },
      },
    ],
    ueAmbr: { downlink: "1 Gbps", uplink: "500 Mbps" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.simPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    mobileNetworkName,
    simPolicyName,
    parameters
  );
  console.log(result);
}

createSimPolicy().catch(console.error);
