
/**
 * Samples for AzureLargeStorageInstance ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeStorageInstance_ListByResourceGroup.json
     */
    /**
     * Sample code: AzureLargeStorageInstance_ListByResourceGroup.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void azureLargeStorageInstanceListByResourceGroup(
        com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeStorageInstances().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
