
/**
 * Samples for DiskEncryptionSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskEncryptionSetExamples/DiskEncryptionSet_Get.json
     */
    /**
     * Sample code: get information about a disk encryption set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutADiskEncryptionSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().getByResourceGroupWithResponse("myResourceGroup",
            "myDiskEncryptionSet", com.azure.core.util.Context.NONE);
    }
}
