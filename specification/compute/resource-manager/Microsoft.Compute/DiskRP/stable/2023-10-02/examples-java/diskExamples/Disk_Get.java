
/**
 * Samples for Disks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskExamples/Disk_Get.
     * json
     */
    /**
     * Sample code: Get information about a managed disk.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().getByResourceGroupWithResponse("myResourceGroup",
            "myManagedDisk", com.azure.core.util.Context.NONE);
    }
}
