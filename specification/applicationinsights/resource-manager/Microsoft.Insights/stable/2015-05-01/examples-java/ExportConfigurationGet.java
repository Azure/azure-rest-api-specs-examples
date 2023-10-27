/** Samples for ExportConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationGet.json
     */
    /**
     * Sample code: ExportConfigurationGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void exportConfigurationGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .exportConfigurations()
            .getWithResponse(
                "my-resource-group", "my-component", "uGOoki0jQsyEs3IdQ83Q4QsNr4=", com.azure.core.util.Context.NONE);
    }
}
