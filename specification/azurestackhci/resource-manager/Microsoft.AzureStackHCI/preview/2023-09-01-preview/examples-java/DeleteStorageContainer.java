/** Samples for StorageContainersOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteStorageContainer.json
     */
    /**
     * Sample code: DeleteStorageContainer.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteStorageContainer(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.storageContainersOperations().delete("test-rg", "Default_Container", com.azure.core.util.Context.NONE);
    }
}
