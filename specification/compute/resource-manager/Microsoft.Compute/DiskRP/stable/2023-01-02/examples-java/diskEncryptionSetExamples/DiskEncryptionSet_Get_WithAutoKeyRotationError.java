/** Samples for DiskEncryptionSets GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Get_WithAutoKeyRotationError.json
     */
    /**
     * Sample code: Get information about a disk encryption set when auto-key rotation failed.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutADiskEncryptionSetWhenAutoKeyRotationFailed(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskEncryptionSets()
            .getByResourceGroupWithResponse("myResourceGroup", "myDiskEncryptionSet", com.azure.core.util.Context.NONE);
    }
}
