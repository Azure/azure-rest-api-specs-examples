
/**
 * Samples for SshPublicKeys GenerateKeyPair.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * sshPublicKeyExamples/SshPublicKey_GenerateKeyPair.json
     */
    /**
     * Sample code: Generate an SSH key pair.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateAnSSHKeyPair(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().generateKeyPairWithResponse(
            "myResourceGroup", "mySshPublicKeyName", null, com.azure.core.util.Context.NONE);
    }
}
