
/**
 * Samples for SshPublicKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * sshPublicKeyExamples/SshPublicKey_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKey_ListBySubscription_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        sshPublicKeyListBySubscriptionMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().list(com.azure.core.util.Context.NONE);
    }
}
