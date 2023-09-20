const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a dedicated host .
 *
 * @summary Update a dedicated host .
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/dedicatedHostExamples/DedicatedHost_Update_MaximumSet_Gen.json
 */
async function dedicatedHostUpdateMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const hostGroupName = "aaaaaaaaa";
  const hostName = "aaaaaaaaaaaaaaaaaaaaa";
  const parameters = {
    autoReplaceOnFailure: true,
    instanceView: {
      availableCapacity: {
        allocatableVMs: [{ count: 26, vmSize: "aaaaaaaaaaaaaaaaaaaa" }],
      },
      statuses: [
        {
          code: "aaaaaaaaaaaaaaaaaaaaaaa",
          displayStatus: "aaaaaa",
          level: "Info",
          message: "a",
          time: new Date("2021-11-30T12:58:26.522Z"),
        },
      ],
    },
    licenseType: "Windows_Server_Hybrid",
    platformFaultDomain: 1,
    tags: { key8813: "aaaaaaaaaaaaaaaaaaaaaaaaaaa" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHosts.beginUpdateAndWait(
    resourceGroupName,
    hostGroupName,
    hostName,
    parameters
  );
  console.log(result);
}
