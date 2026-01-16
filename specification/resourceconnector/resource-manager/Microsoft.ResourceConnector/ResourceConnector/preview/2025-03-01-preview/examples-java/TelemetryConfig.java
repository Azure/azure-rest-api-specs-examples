
/**
 * Samples for Appliances GetTelemetryConfig.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/TelemetryConfig.json
     */
    /**
     * Sample code: GetTelemetryConfig Appliance.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void
        getTelemetryConfigAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().getTelemetryConfigWithResponse(com.azure.core.util.Context.NONE);
    }
}
