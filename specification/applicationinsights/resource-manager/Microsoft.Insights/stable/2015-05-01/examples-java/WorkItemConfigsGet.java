/** Samples for WorkItemConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigsGet.json
     */
    /**
     * Sample code: WorkItemConfigurationsList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workItemConfigurationsList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.workItemConfigurations().list("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
