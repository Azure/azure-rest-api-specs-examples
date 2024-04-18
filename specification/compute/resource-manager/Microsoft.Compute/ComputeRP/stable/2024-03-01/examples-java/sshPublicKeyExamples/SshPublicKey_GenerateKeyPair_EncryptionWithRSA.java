
import com.azure.resourcemanager.compute.models.SshEncryptionTypes;
import com.azure.resourcemanager.compute.models.SshGenerateKeyPairInputParameters;

/**
 * Samples for SshPublicKeys GenerateKeyPair.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * sshPublicKeyExamples/SshPublicKey_GenerateKeyPair_EncryptionWithRSA.json
     */
    /**
     * Sample code: Generate an SSH key pair with RSA encryption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateAnSSHKeyPairWithRSAEncryption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().generateKeyPairWithResponse(
            "myResourceGroup", "mySshPublicKeyName",
            new SshGenerateKeyPairInputParameters().withEncryptionType(SshEncryptionTypes.RSA),
            com.azure.core.util.Context.NONE);
    }
}
