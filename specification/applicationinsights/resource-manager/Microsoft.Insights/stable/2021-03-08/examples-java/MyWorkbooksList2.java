import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for MyWorkbooks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbooksList2.json
     */
    /**
     * Sample code: WorkbooksList2.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbooksList2(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .myWorkbooks()
            .listByResourceGroup(
                "my-resource-group", CategoryType.WORKBOOK, null, null, null, com.azure.core.util.Context.NONE);
    }
}
