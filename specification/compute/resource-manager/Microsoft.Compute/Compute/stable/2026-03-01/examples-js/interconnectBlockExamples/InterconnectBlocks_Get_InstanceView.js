const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about an Interconnect Block.
 *
 * @summary retrieves information about an Interconnect Block.
 * x-ms-original-file: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_Get_InstanceView.json
 */
async function getAnInterconnectBlockWithInstanceView() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.interconnectBlocks.get("myResourceGroup", "myInterconnectBlock", {
    expand: "instanceView",
  });
  console.log(result);
}
