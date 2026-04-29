
/**
 * Samples for DiskAccesses List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskAccessExamples/DiskAccess_ListBySubscription.json
     */
    /**
     * Sample code: list all disk access resources in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listAllDiskAccessResourcesInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().list(com.azure.core.util.Context.NONE);
    }
}
