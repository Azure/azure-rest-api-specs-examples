
/**
 * Samples for DiskAccesses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskAccessExamples/
     * DiskAccess_Delete.json
     */
    /**
     * Sample code: Delete a disk access resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteADiskAccessResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskAccesses().delete("myResourceGroup", "myDiskAccess",
            com.azure.core.util.Context.NONE);
    }
}
