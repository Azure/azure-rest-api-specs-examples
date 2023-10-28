/** Samples for WorkbookTemplates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateDelete.json
     */
    /**
     * Sample code: WorkbookTemplateDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookTemplateDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbookTemplates()
            .deleteByResourceGroupWithResponse(
                "my-resource-group", "my-template-resource", com.azure.core.util.Context.NONE);
    }
}
