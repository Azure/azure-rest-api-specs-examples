
/**
 * Samples for SshPublicKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * sshPublicKeyExamples/SshPublicKey_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_Delete_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sshPublicKeyDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().deleteWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
