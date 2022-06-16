import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for MyWorkbooks List. */
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
        manager.myWorkbooks().list(CategoryType.WORKBOOK, null, null, Context.NONE);
    }
}
