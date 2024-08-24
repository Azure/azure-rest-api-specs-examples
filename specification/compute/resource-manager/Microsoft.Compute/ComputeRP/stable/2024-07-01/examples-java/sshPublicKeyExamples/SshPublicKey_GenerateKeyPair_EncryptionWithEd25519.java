
import com.azure.resourcemanager.compute.models.SshEncryptionTypes;
import com.azure.resourcemanager.compute.models.SshGenerateKeyPairInputParameters;

/**
 * Samples for SshPublicKeys GenerateKeyPair.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * sshPublicKeyExamples/SshPublicKey_GenerateKeyPair_EncryptionWithEd25519.json
     */
    /**
     * Sample code: Generate an SSH key pair with Ed25519 encryption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateAnSSHKeyPairWithEd25519Encryption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().generateKeyPairWithResponse(
            "myResourceGroup", "mySshPublicKeyName",
            new SshGenerateKeyPairInputParameters().withEncryptionType(SshEncryptionTypes.ED25519),
            com.azure.core.util.Context.NONE);
    }
}
