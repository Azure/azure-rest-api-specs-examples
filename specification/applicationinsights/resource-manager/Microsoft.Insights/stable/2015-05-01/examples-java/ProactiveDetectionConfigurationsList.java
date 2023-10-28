/** Samples for ProactiveDetectionConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationsList.json
     */
    /**
     * Sample code: ProactiveDetectionConfigurationsList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void proactiveDetectionConfigurationsList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .proactiveDetectionConfigurations()
            .listWithResponse("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
