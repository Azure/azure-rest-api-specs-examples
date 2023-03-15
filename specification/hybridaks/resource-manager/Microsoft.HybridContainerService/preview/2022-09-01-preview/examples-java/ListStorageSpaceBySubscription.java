/** Samples for StorageSpacesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ListStorageSpaceBySubscription.json
     */
    /**
     * Sample code: ListStorageSpaceBySubscription.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listStorageSpaceBySubscription(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.storageSpacesOperations().list(com.azure.core.util.Context.NONE);
    }
}
