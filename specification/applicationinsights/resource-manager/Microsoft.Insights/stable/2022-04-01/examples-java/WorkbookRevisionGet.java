
/**
 * Samples for Workbooks RevisionGet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/
     * WorkbookRevisionGet.json
     */
    /**
     * Sample code: WorkbookRevisionGet.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void
        workbookRevisionGet(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.workbooks().revisionGetWithResponse("my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
            "1e2f8435b98248febee70c64ac22e1ab", com.azure.core.util.Context.NONE);
    }
}
