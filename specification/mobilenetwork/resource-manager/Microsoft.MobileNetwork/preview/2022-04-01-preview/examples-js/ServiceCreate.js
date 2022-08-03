const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a service.
 *
 * @summary Creates or updates a service.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/ServiceCreate.json
 */
async function createService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const serviceName = "TestService";
  const parameters = {
    location: "eastus",
    pccRules: [
      {
        ruleName: "default-rule",
        rulePrecedence: 255,
        ruleQosPolicy: {
          fiveQi: 9,
          allocationAndRetentionPriorityLevel: 9,
          maximumBitRate: { downlink: "1 Gbps", uplink: "500 Mbps" },
          preemptionCapability: "NotPreempt",
          preemptionVulnerability: "Preemptable",
        },
        serviceDataFlowTemplates: [
          {
            direction: "Uplink",
            ports: [],
            remoteIpList: ["10.3.4.0/24"],
            templateName: "IP-to-server",
            protocol: ["ip"],
          },
        ],
        trafficControl: "Enabled",
      },
    ],
    servicePrecedence: 255,
    serviceQosPolicy: {
      fiveQi: 9,
      allocationAndRetentionPriorityLevel: 9,
      maximumBitRate: { downlink: "1 Gbps", uplink: "500 Mbps" },
      preemptionCapability: "NotPreempt",
      preemptionVulnerability: "Preemptable",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.services.beginCreateOrUpdateAndWait(
    resourceGroupName,
    mobileNetworkName,
    serviceName,
    parameters
  );
  console.log(result);
}

createService().catch(console.error);
