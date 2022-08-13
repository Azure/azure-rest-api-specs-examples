import com.azure.core.util.Context;

/** Samples for NodeReports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeReport.json
     */
    /**
     * Sample code: Get Dsc node report data by node id and report id.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getDscNodeReportDataByNodeIdAndReportId(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .nodeReports()
            .getWithResponse(
                "rg", "myAutomationAccount33", "nodeId", "903a5ead-140c-11e7-a943-000d3a6140c9", Context.NONE);
    }
}
