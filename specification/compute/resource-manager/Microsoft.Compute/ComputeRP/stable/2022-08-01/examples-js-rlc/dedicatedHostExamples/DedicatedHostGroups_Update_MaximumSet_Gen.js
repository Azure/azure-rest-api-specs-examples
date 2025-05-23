const createComputeManagementClient = require("@azure-rest/arm-compute").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Update an dedicated host group.
 *
 * @summary Update an dedicated host group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/dedicatedHostExamples/DedicatedHostGroups_Update_MaximumSet_Gen.json
 */
async function dedicatedHostGroupsUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rgcompute";
  const hostGroupName = "aaaa";
  const options = {
    body: {
      properties: {
        platformFaultDomainCount: 3,
        supportAutomaticPlacement: true,
      },
      tags: { key9921: "aaaaaaaaaa" },
      zones: ["aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"],
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}",
      subscriptionId,
      resourceGroupName,
      hostGroupName,
    )
    .patch(options);
  console.log(result);
}

dedicatedHostGroupsUpdateMaximumSetGen().catch(console.error);
