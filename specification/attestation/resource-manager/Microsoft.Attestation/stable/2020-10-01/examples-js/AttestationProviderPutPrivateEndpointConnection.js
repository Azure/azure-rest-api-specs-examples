const { AttestationManagementClient } = require("@azure/arm-attestation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified private endpoint connection associated with the attestation provider.
 *
 * @summary Update the state of specified private endpoint connection associated with the attestation provider.
 * x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderPutPrivateEndpointConnection.json
 */
async function attestationProviderPutPrivateEndpointConnection() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const providerName = "sto9699";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AttestationManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.create(
    resourceGroupName,
    providerName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}

attestationProviderPutPrivateEndpointConnection().catch(console.error);
