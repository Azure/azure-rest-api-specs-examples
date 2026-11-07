
/**
 * Samples for AutoscaleSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/listAutoscaleSettingBySubscription.json
     */
    /**
     * Sample code: List autoscale settings by subs.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listAutoscaleSettingsBySubs(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getAutoscaleSettings().list(com.azure.core.util.Context.NONE);
    }
}
