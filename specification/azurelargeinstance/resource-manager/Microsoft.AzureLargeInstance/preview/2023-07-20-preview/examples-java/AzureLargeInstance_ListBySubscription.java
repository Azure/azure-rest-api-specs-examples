
/**
 * Samples for AzureLargeInstance List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_ListBySubscription.json
     */
    /**
     * Sample code: AzureLargeInstance_ListBySubscription.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeInstanceListBySubscription(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().list(com.azure.core.util.Context.NONE);
    }
}
