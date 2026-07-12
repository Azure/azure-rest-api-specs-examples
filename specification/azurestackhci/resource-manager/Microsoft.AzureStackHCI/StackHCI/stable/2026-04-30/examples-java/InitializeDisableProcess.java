
/**
 * Samples for ArcSettings InitializeDisableProcess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/InitializeDisableProcess.json
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
