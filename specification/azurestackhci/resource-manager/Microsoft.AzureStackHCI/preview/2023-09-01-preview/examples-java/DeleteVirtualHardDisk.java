/** Samples for VirtualHardDisksOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteVirtualHardDisk.json
     */
    /**
     * Sample code: DeleteVirtualHardDisk.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteVirtualHardDisk(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.virtualHardDisksOperations().delete("test-rg", "test-vhd", com.azure.core.util.Context.NONE);
    }
}
