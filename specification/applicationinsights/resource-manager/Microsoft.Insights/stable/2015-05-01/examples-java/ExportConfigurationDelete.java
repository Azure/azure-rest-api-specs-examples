/** Samples for ExportConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationDelete.json
     */
    /**
     * Sample code: ExportConfigurationDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void exportConfigurationDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .exportConfigurations()
            .deleteWithResponse(
                "my-resource-group", "my-component", "uGOoki0jQsyEs3IdQ83Q4QsNr4=", com.azure.core.util.Context.NONE);
    }
}
