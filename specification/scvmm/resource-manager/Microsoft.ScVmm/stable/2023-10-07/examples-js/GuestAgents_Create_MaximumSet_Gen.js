const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update GuestAgent.
 *
 * @summary Create Or Update GuestAgent.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GuestAgents_Create_MaximumSet_Gen.json
 */
async function guestAgentsCreateMaximumSet() {
  const resourceUri = "gtgclehcbsyave";
  const resource = {
    properties: {
      credentials: {
        password: "gkvbnmuahumuoibvscoxzfdqwvfuf",
        username: "jqxuwirrcpfv",
      },
      httpProxyConfig: { httpsProxy: "uoyzyticmohohomlkwct" },
      provisioningAction: "install",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.guestAgents.beginCreateAndWait(resourceUri, resource);
  console.log(result);
}
