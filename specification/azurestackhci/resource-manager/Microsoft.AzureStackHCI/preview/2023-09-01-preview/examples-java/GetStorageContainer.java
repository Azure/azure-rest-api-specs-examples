/** Samples for StorageContainersOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetStorageContainer.json
     */
    /**
     * Sample code: GetStorageContainer.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getStorageContainer(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .storageContainersOperations()
            .getByResourceGroupWithResponse("test-rg", "Default_Container", com.azure.core.util.Context.NONE);
    }
}
