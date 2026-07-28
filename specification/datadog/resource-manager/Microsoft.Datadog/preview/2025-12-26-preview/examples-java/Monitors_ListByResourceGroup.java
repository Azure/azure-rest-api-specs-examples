
/**
 * Samples for Monitors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_ListByResourceGroup.json
     */
    /**
     * Sample code: Monitors_ListByResourceGroup.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsListByResourceGroup(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
