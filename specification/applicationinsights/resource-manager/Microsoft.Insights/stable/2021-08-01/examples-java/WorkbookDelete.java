import com.azure.core.util.Context;

/** Samples for Workbooks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookDelete.json
     */
    /**
     * Sample code: WorkbookDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .deleteWithResponse("my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", Context.NONE);
    }
}
