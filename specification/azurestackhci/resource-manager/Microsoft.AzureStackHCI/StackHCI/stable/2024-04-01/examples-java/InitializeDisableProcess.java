
/**
 * Samples for ArcSettings InitializeDisableProcess.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * InitializeDisableProcess.json
     */
    /**
     * Sample code: Trigger ARC Disable.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void triggerARCDisable(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().initializeDisableProcess("test-rg", "myCluster", "default",
            com.azure.core.util.Context.NONE);
    }
}
