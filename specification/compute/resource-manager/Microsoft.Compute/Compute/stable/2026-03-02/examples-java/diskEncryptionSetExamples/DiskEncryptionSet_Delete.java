
/**
 * Samples for DiskEncryptionSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskEncryptionSetExamples/DiskEncryptionSet_Delete.json
     */
    /**
     * Sample code: delete a disk encryption set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteADiskEncryptionSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().delete("myResourceGroup", "myDiskEncryptionSet",
            com.azure.core.util.Context.NONE);
    }
}
