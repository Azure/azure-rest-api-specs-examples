const { AttestationManagementClient } = require("@azure/arm-attestation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private endpoint connection associated with the attestation provider.
 *
 * @summary Deletes the specified private endpoint connection associated with the attestation provider.
 * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderDeletePrivateEndpointConnection.json
 */
async function attestationProviderDeletePrivateEndpointConnection() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const providerName = "sto2527";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const credential = new DefaultAzureCredential();
  const client = new AttestationManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    providerName,
    privateEndpointConnectionName
  );
  console.log(result);
}

attestationProviderDeletePrivateEndpointConnection().catch(console.error);
