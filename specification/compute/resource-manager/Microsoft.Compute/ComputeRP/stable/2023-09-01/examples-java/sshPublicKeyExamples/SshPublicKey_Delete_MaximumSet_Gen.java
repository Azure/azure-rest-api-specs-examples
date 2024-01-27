
/**
 * Samples for SshPublicKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * sshPublicKeyExamples/SshPublicKey_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_Delete_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sshPublicKeyDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().deleteWithResponse("rgcompute",
            "aaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
