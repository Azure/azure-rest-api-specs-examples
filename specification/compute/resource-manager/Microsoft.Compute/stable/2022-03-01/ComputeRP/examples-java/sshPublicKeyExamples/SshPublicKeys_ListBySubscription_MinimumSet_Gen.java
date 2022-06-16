import com.azure.core.util.Context;

/** Samples for SshPublicKeys List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/sshPublicKeyExamples/SshPublicKeys_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: SshPublicKeys_ListBySubscription_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sshPublicKeysListBySubscriptionMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSshPublicKeys().list(Context.NONE);
    }
}
