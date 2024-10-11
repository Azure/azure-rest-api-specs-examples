
/**
 * Samples for NodeReports ListByNode.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/
     * listPagedDscNodeReportsByNode.json
     */
    /**
     * Sample code: List Paged DSC reports by node id.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCReportsByNodeId(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.nodeReports().listByNode("rg", "myAutomationAccount33", "nodeId", null,
            com.azure.core.util.Context.NONE);
    }
}
