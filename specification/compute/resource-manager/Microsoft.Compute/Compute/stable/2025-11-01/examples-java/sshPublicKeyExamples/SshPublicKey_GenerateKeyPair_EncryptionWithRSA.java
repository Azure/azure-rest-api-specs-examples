
import com.azure.resourcemanager.compute.models.SshEncryptionTypes;
import com.azure.resourcemanager.compute.models.SshGenerateKeyPairInputParameters;

/**
 * Samples for SshPublicKeys GenerateKeyPair.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_GenerateKeyPair_EncryptionWithRSA.json
     */
    /**
     * Sample code: Generate an SSH key pair with RSA encryption.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void generateAnSSHKeyPairWithRSAEncryption(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().generateKeyPairWithResponse("myResourceGroup", "mySshPublicKeyName",
            new SshGenerateKeyPairInputParameters().withEncryptionType(SshEncryptionTypes.RSA),
            com.azure.core.util.Context.NONE);
    }
}
