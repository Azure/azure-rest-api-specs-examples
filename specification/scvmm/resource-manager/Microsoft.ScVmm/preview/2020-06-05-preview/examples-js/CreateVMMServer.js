const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the SCVMM fabric as an Azure VmmServer resource.
 *
 * @summary Onboards the SCVMM fabric as an Azure VmmServer resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateVMMServer.json
 */
async function createVmmServer() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const vmmServerName = "ContosoVMMServer";
  const body = {
    credentials: { password: "password", username: "testuser" },
    extendedLocation: {
      name: "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso",
      type: "customLocation",
    },
    fqdn: "VMM.contoso.com",
    location: "East US",
    port: 1234,
  };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.vmmServers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmmServerName,
    body
  );
  console.log(result);
}
