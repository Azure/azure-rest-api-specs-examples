
/**
 * Samples for AzureLargeInstance ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_ListByResourceGroup.json
     */
    /**
     * Sample code: AzureLargeInstance_ListByResourceGroup.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeInstanceListByResourceGroup(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
