const { AttestationManagementClient } = require("@azure/arm-attestation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the private endpoint connections associated with the attestation provider.
 *
 * @summary List all the private endpoint connections associated with the attestation provider.
 * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderListPrivateEndpointConnections.json
 */
async function attestationProviderListPrivateEndpointConnections() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const providerName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new AttestationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, providerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

attestationProviderListPrivateEndpointConnections().catch(console.error);
