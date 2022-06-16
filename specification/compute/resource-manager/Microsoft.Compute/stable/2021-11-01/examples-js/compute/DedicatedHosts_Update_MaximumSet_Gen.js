const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function dedicatedHostsUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
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

dedicatedHostsUpdateMaximumSetGen().catch(console.error);
