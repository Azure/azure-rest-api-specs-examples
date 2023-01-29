/** Samples for SshPublicKeys ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/sshPublicKeyExamples/SshPublicKeys_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKeys_ListByResourceGroup_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sshPublicKeysListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSshPublicKeys()
            .listByResourceGroup("rgcompute", com.azure.core.util.Context.NONE);
    }
}
