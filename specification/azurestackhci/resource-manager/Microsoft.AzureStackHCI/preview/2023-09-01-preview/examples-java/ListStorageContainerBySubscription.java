/** Samples for StorageContainersOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListStorageContainerBySubscription.json
     */
    /**
     * Sample code: ListStorageContainerBySubscription.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listStorageContainerBySubscription(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.storageContainersOperations().list(com.azure.core.util.Context.NONE);
    }
}
