import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for Workbooks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbooksList.json
     */
    /**
     * Sample code: WorkbooksList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbooksList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .listByResourceGroup(
                "my-resource-group",
                CategoryType.WORKBOOK,
                null,
                "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/my-resource-group/providers/Microsoft.Web/sites/MyApp",
                null,
                com.azure.core.util.Context.NONE);
    }
}
