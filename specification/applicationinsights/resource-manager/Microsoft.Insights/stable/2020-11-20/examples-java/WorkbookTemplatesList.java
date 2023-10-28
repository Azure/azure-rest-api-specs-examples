/** Samples for WorkbookTemplates ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplatesList.json
     */
    /**
     * Sample code: WorkbookTemplatesList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookTemplatesList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.workbookTemplates().listByResourceGroup("my-resource-group", com.azure.core.util.Context.NONE);
    }
}
