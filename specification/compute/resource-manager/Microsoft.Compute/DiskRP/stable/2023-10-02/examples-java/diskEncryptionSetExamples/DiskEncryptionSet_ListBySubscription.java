
/**
 * Samples for DiskEncryptionSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/
     * diskEncryptionSetExamples/DiskEncryptionSet_ListBySubscription.json
     */
    /**
     * Sample code: List all disk encryption sets in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllDiskEncryptionSetsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskEncryptionSets()
            .list(com.azure.core.util.Context.NONE);
    }
}
