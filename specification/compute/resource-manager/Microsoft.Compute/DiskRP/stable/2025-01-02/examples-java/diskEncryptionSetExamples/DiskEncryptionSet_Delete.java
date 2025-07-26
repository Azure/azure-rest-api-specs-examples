
/**
 * Samples for DiskEncryptionSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/
     * diskEncryptionSetExamples/DiskEncryptionSet_Delete.json
     */
    /**
     * Sample code: delete a disk encryption set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteADiskEncryptionSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskEncryptionSets().delete("myResourceGroup",
            "myDiskEncryptionSet", com.azure.core.util.Context.NONE);
    }
}
