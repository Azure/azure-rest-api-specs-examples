
/**
 * Samples for DiskEncryptionSets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskEncryptionSetExamples/DiskEncryptionSet_ListByResourceGroup.json
     */
    /**
     * Sample code: list all disk encryption sets in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listAllDiskEncryptionSetsInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
