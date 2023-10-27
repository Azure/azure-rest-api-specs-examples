import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for MyWorkbooks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbooksList.json
     */
    /**
     * Sample code: WorkbooksList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbooksList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .myWorkbooks()
            .listByResourceGroup(
                "my-resource-group", CategoryType.WORKBOOK, null, null, null, com.azure.core.util.Context.NONE);
    }
}
