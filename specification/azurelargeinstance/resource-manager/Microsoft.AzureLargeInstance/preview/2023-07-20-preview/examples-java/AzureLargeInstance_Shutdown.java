
/**
 * Samples for AzureLargeInstance Shutdown.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_Shutdown.json
     */
    /**
     * Sample code: AzureLargeInstance_Shutdown.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeInstanceShutdown(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().shutdown("myResourceGroup", "myALInstance", com.azure.core.util.Context.NONE);
    }
}
