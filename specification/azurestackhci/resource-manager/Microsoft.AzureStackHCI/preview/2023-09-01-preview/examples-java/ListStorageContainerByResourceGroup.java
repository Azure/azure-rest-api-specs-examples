/** Samples for StorageContainersOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListStorageContainerByResourceGroup.json
     */
    /**
     * Sample code: ListStorageContainerByResourceGroup.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listStorageContainerByResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.storageContainersOperations().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
