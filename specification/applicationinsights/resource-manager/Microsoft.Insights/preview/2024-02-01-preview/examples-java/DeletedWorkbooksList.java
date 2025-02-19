
/**
 * Samples for DeletedWorkbooks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2024-02-01-preview/examples/
     * DeletedWorkbooksList.json
     */
    /**
     * Sample code: WorkbooksListSub.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void
        workbooksListSub(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.deletedWorkbooks().list(null, null, com.azure.core.util.Context.NONE);
    }
}
