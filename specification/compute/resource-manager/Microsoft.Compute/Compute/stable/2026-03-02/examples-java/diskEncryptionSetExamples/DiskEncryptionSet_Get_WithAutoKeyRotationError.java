
/**
 * Samples for DiskEncryptionSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskEncryptionSetExamples/DiskEncryptionSet_Get_WithAutoKeyRotationError.json
     */
    /**
     * Sample code: get information about a disk encryption set when auto-key rotation failed.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutADiskEncryptionSetWhenAutoKeyRotationFailed(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().getByResourceGroupWithResponse("myResourceGroup",
            "myDiskEncryptionSet", com.azure.core.util.Context.NONE);
    }
}
