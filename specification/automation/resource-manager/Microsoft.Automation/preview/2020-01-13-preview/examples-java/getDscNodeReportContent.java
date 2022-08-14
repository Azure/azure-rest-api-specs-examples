import com.azure.core.util.Context;

/** Samples for NodeReports GetContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeReportContent.json
     */
    /**
     * Sample code: Get content of node.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getContentOfNode(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.nodeReports().getContentWithResponse("rg", "myAutomationAccount33", "nodeId", "reportId", Context.NONE);
    }
}
