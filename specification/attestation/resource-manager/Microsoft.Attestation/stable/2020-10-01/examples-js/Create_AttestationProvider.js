const { AttestationManagementClient } = require("@azure/arm-attestation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Attestation Provider.
 *
 * @summary Creates a new Attestation Provider.
 * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Create_AttestationProvider.json
 */
async function attestationProvidersCreate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "MyResourceGroup";
  const providerName = "myattestationprovider";
  const creationParams = { location: "eastus", properties: {} };
  const credential = new DefaultAzureCredential();
  const client = new AttestationManagementClient(credential, subscriptionId);
  const result = await client.attestationProviders.create(
    resourceGroupName,
    providerName,
    creationParams
  );
  console.log(result);
}

attestationProvidersCreate().catch(console.error);
