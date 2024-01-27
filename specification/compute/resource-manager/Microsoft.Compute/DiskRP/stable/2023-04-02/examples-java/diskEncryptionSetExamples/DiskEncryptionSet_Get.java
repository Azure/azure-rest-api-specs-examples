
/**
 * Samples for DiskEncryptionSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/
     * diskEncryptionSetExamples/DiskEncryptionSet_Get.json
     */
    /**
     * Sample code: Get information about a disk encryption set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutADiskEncryptionSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskEncryptionSets()
            .getByResourceGroupWithResponse("myResourceGroup", "myDiskEncryptionSet", com.azure.core.util.Context.NONE);
    }
}
