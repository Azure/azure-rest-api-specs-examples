const { AttestationManagementClient } = require("@azure/arm-attestation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Attestation Provider.
 *
 * @summary Updates the Attestation Provider.
 * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Update_AttestationProvider.json
 */
async function attestationProvidersUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "MyResourceGroup";
  const providerName = "myattestationprovider";
  const updateParams = {
    tags: { property1: "Value1", property2: "Value2", property3: "Value3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AttestationManagementClient(credential, subscriptionId);
  const result = await client.attestationProviders.update(
    resourceGroupName,
    providerName,
    updateParams
  );
  console.log(result);
}

attestationProvidersUpdate().catch(console.error);
