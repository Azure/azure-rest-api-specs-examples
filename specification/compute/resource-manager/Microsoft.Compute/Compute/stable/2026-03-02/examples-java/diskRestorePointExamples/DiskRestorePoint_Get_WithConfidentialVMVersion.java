
/**
 * Samples for DiskRestorePoint Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_Get_WithConfidentialVMVersion.json
     */
    /**
     * Sample code: get a ConfidentialVM incremental disk restorePoint resource with confidentialVMVersion.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAConfidentialVMIncrementalDiskRestorePointResourceWithConfidentialVMVersion(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskRestorePoints().getWithResponse("myResourceGroup", "rpc", "vmrp",
            "myConfidentialDisk_c4bc27e0-ccf6-494e-a740-af34de775527", com.azure.core.util.Context.NONE);
    }
}
