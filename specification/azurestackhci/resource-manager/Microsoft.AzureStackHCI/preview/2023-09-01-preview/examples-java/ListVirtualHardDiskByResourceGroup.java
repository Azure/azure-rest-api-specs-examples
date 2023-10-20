/** Samples for VirtualHardDisksOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListVirtualHardDiskByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualHardDiskByResourceGroup.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listVirtualHardDiskByResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.virtualHardDisksOperations().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
