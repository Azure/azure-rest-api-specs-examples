
/**
 * Samples for AutoscaleSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/deleteAutoscaleSetting.json
     */
    /**
     * Sample code: Delete an autoscale setting.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteAnAutoscaleSetting(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getAutoscaleSettings().deleteWithResponse("TestingMetricsScaleSet", "MySetting",
            com.azure.core.util.Context.NONE);
    }
}
