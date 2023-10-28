import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for Workbooks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbooksListSub.json
     */
    /**
     * Sample code: WorkbooksListSub.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbooksListSub(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.workbooks().list(CategoryType.WORKBOOK, null, null, com.azure.core.util.Context.NONE);
    }
}
