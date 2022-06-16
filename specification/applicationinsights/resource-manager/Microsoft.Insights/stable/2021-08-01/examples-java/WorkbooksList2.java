import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for Workbooks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbooksList2.json
     */
    /**
     * Sample code: WorkbooksList2.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbooksList2(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.workbooks().list(CategoryType.WORKBOOK, null, null, Context.NONE);
    }
}
