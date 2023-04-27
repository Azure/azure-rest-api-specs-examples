/** Samples for Appliances GetTelemetryConfig. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/TelemetryConfig.json
     */
    /**
     * Sample code: GetTelemetryConfig Appliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void getTelemetryConfigAppliance(
        com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().getTelemetryConfigWithResponse(com.azure.core.util.Context.NONE);
    }
}
