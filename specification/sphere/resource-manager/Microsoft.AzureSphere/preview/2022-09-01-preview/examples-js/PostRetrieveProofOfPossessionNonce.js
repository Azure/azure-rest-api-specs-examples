const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the proof of possession nonce.
 *
 * @summary Gets the proof of possession nonce.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostRetrieveProofOfPossessionNonce.json
 */
async function certificatesRetrieveProofOfPossessionNonce() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const serialNumber = "active";
  const proofOfPossessionNonceRequest = {
    proofOfPossessionNonce: "proofOfPossessionNonce",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const result = await client.certificates.retrieveProofOfPossessionNonce(
    resourceGroupName,
    catalogName,
    serialNumber,
    proofOfPossessionNonceRequest
  );
  console.log(result);
}
