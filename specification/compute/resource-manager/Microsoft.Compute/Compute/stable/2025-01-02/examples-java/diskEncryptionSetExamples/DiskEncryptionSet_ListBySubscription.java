
/**
 * Samples for DiskEncryptionSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskEncryptionSetExamples/DiskEncryptionSet_ListBySubscription.json
     */
    /**
     * Sample code: list all disk encryption sets in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listAllDiskEncryptionSetsInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().list(com.azure.core.util.Context.NONE);
    }
}
