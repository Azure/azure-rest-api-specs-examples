/** Samples for VirtualHardDisksOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListVirtualHardDiskBySubscription.json
     */
    /**
     * Sample code: ListVirtualHardDiskBySubscription.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listVirtualHardDiskBySubscription(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.virtualHardDisksOperations().list(com.azure.core.util.Context.NONE);
    }
}
