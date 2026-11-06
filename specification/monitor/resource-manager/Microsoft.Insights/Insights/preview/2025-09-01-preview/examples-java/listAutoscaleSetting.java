
/**
 * Samples for AutoscaleSettings ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/listAutoscaleSetting.json
     */
    /**
     * Sample code: List autoscale settings.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listAutoscaleSettings(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getAutoscaleSettings().listByResourceGroup("TestingMetricsScaleSet",
            com.azure.core.util.Context.NONE);
    }
}
