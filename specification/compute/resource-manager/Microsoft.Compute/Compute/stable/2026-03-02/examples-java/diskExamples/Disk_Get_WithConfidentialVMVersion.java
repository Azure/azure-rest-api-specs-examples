
/**
 * Samples for Disks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Get_WithConfidentialVMVersion.json
     */
    /**
     * Sample code: get information about a ConfidentialVM disk with confidentialVMVersion.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutAConfidentialVMDiskWithConfidentialVMVersion(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().getByResourceGroupWithResponse("myResourceGroup", "myConfidentialDisk",
            com.azure.core.util.Context.NONE);
    }
}
