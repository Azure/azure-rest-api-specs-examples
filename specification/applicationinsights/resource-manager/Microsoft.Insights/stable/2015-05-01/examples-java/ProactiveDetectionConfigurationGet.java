/** Samples for ProactiveDetectionConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationGet.json
     */
    /**
     * Sample code: ProactiveDetectionConfigurationGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void proactiveDetectionConfigurationGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .proactiveDetectionConfigurations()
            .getWithResponse("my-resource-group", "my-component", "slowpageloadtime", com.azure.core.util.Context.NONE);
    }
}
