const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to generates and returns a public/private key pair and populates the SSH public key resource with the public key. The length of the key will be 3072 bits. This operation can only be performed once per SSH public key resource.
 *
 * @summary generates and returns a public/private key pair and populates the SSH public key resource with the public key. The length of the key will be 3072 bits. This operation can only be performed once per SSH public key resource.
 * x-ms-original-file: 2026-03-01/sshPublicKeyExamples/SshPublicKey_GenerateKeyPair_EncryptionWithEd25519.json
 */
async function generateAnSSHKeyPairWithEd25519Encryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.generateKeyPair(
    "myResourceGroup",
    "mySshPublicKeyName",
    { parameters: { encryptionType: "Ed25519" } },
  );
  console.log(result);
}
