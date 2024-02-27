
/**
 * Samples for AzureLargeStorageInstance List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeStorageInstance_ListBySubscription.json
     */
    /**
     * Sample code: AzureLargeStorageInstance_ListBySubscription.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void azureLargeStorageInstanceListBySubscription(
        com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeStorageInstances().list(com.azure.core.util.Context.NONE);
    }
}
