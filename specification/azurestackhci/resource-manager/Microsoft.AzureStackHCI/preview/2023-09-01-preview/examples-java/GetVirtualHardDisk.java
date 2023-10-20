/** Samples for VirtualHardDisksOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetVirtualHardDisk.json
     */
    /**
     * Sample code: GetVirtualHardDisk.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getVirtualHardDisk(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .virtualHardDisksOperations()
            .getByResourceGroupWithResponse("test-rg", "test-vhd", com.azure.core.util.Context.NONE);
    }
}
