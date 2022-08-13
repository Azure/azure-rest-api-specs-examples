import com.azure.core.util.Context;

/** Samples for NodeReports ListByNode. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listDscNodeReportsByNode.json
     */
    /**
     * Sample code: List DSC reports by node id.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listDSCReportsByNodeId(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.nodeReports().listByNode("rg", "myAutomationAccount33", "nodeId", null, Context.NONE);
    }
}
