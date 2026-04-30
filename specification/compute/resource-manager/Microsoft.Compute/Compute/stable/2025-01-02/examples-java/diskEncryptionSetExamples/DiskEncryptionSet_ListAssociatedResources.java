
/**
 * Samples for DiskEncryptionSets ListAssociatedResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskEncryptionSetExamples/DiskEncryptionSet_ListAssociatedResources.json
     */
    /**
     * Sample code: list all resources that are encrypted with this disk encryption set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAllResourcesThatAreEncryptedWithThisDiskEncryptionSet(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().listAssociatedResources("myResourceGroup",
            "myDiskEncryptionSet", com.azure.core.util.Context.NONE);
    }
}
