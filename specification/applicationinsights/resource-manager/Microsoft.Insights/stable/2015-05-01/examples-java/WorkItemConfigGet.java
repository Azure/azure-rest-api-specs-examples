/** Samples for WorkItemConfigurations GetItem. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigGet.json
     */
    /**
     * Sample code: WorkItemConfigurationsGetDefault.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workItemConfigurationsGetDefault(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workItemConfigurations()
            .getItemWithResponse(
                "my-resource-group", "my-component", "Visual Studio Team Services", com.azure.core.util.Context.NONE);
    }
}
