
/**
 * Samples for DiskEncryptionSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/
     * diskEncryptionSetExamples/DiskEncryptionSet_ListBySubscription.json
     */
    /**
     * Sample code: list all disk encryption sets in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllDiskEncryptionSetsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskEncryptionSets()
            .list(com.azure.core.util.Context.NONE);
    }
}
