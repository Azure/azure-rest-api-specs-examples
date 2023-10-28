/** Samples for WorkbookTemplates GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateGet.json
     */
    /**
     * Sample code: WorkbookTemplateGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookTemplateGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbookTemplates()
            .getByResourceGroupWithResponse("my-resource-group", "my-resource-name", com.azure.core.util.Context.NONE);
    }
}
