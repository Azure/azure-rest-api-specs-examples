
/**
 * Samples for Disks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_ListBySubscription.json
     */
    /**
     * Sample code: list all managed disks in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAllManagedDisksInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().list(com.azure.core.util.Context.NONE);
    }
}
