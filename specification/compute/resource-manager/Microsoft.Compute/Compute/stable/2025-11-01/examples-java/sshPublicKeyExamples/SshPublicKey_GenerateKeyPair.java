
/**
 * Samples for SshPublicKeys GenerateKeyPair.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_GenerateKeyPair.json
     */
    /**
     * Sample code: Generate an SSH key pair.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void generateAnSSHKeyPair(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSshPublicKeys().generateKeyPairWithResponse("myResourceGroup", "mySshPublicKeyName",
            null, com.azure.core.util.Context.NONE);
    }
}
