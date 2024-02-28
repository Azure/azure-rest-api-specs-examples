
/**
 * Samples for AzureLargeInstance Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_Restart.json
     */
    /**
     * Sample code: AzureLargeInstance_Restart.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void azureLargeInstanceRestart(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().restart("myResourceGroup", "myALInstance", null,
            com.azure.core.util.Context.NONE);
    }
}
