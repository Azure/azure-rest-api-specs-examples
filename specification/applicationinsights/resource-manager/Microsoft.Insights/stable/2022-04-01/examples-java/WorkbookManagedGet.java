/** Samples for Workbooks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookManagedGet.json
     */
    /**
     * Sample code: WorkbookManagedGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookManagedGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .getByResourceGroupWithResponse(
                "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", null, com.azure.core.util.Context.NONE);
    }
}
