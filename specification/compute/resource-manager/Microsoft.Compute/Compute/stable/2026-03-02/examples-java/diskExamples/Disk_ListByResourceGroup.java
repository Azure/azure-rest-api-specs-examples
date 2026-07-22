
/**
 * Samples for Disks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_ListByResourceGroup.json
     */
    /**
     * Sample code: list all managed disks in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAllManagedDisksInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
