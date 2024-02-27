
/**
 * Samples for AzureLargeInstance Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_Start.json
     */
    /**
     * Sample code: AzureLargeInstance_Start.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void azureLargeInstanceStart(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().start("myResourceGroup", "myALInstance", com.azure.core.util.Context.NONE);
    }
}
