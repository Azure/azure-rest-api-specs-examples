
/**
 * Samples for Schema Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SavedSearchesGetSchema.json
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
