
/**
 * Samples for MyWorkbooks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookGet.
     * json
     */
    /**
     * Sample code: WorkbookGet.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookGet(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.myWorkbooks().getByResourceGroupWithResponse("my-resource-group",
            "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", com.azure.core.util.Context.NONE);
    }
}
