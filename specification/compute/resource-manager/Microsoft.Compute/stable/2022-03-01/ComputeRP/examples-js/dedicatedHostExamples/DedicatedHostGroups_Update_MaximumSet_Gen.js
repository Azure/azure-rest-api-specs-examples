const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an dedicated host group.
 *
 * @summary Update an dedicated host group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroups_Update_MaximumSet_Gen.json
 */
async function dedicatedHostGroupsUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const hostGroupName = "aaaa";
  const parameters = {
    instanceView: {
      hosts: [
        {
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
      ],
    },
    platformFaultDomainCount: 3,
    supportAutomaticPlacement: true,
    tags: { key9921: "aaaaaaaaaa" },
    zones: ["aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.update(
    resourceGroupName,
    hostGroupName,
    parameters
  );
  console.log(result);
}

dedicatedHostGroupsUpdateMaximumSetGen().catch(console.error);
