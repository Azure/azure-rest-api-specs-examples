const { AttestationManagementClient } = require("@azure/arm-attestation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns attestation providers list in a resource group.
 *
 * @summary Returns attestation providers list in a resource group.
 * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvidersListByResourceGroup.json
 */
async function attestationProvidersListByResourceGroup() {
  const subscriptionId = "6c96b33e-f5b8-40a6-9011-5cb1c58b0915";
  const resourceGroupName = "testrg1";
  const credential = new DefaultAzureCredential();
  const client = new AttestationManagementClient(credential, subscriptionId);
  const result = await client.attestationProviders.listByResourceGroup(resourceGroupName);
  console.log(result);
}

attestationProvidersListByResourceGroup().catch(console.error);
