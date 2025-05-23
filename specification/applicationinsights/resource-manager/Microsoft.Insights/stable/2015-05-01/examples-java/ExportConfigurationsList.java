
/**
 * Samples for ExportConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/
     * ExportConfigurationsList.json
     */
    /**
     * Sample code: ExportConfigurationsList.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void
        exportConfigurationsList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.exportConfigurations().listWithResponse("my-resource-group", "my-component",
            com.azure.core.util.Context.NONE);
    }
}
