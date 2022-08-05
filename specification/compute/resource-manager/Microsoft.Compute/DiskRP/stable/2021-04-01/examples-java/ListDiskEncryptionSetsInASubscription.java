import com.azure.core.util.Context;

/** Samples for DiskEncryptionSets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/ListDiskEncryptionSetsInASubscription.json
     */
    /**
     * Sample code: List all disk encryption sets in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllDiskEncryptionSetsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskEncryptionSets().list(Context.NONE);
    }
}
