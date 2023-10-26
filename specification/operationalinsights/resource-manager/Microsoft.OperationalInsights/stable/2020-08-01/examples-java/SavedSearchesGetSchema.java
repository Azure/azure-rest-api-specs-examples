/** Samples for Schema Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/SavedSearchesGetSchema.json
     */
    /**
     * Sample code: WorkspacesGetSchema.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesGetSchema(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.schemas().getWithResponse("mms-eus", "atlantisdemo", com.azure.core.util.Context.NONE);
    }
}
