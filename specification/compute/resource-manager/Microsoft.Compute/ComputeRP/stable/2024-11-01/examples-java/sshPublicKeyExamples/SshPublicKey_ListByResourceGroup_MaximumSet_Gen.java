
/**
 * Samples for SshPublicKeys ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        sshPublicKeyListByResourceGroupMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
