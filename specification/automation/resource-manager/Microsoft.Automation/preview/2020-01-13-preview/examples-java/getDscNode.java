import com.azure.core.util.Context;

/** Samples for DscNode Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNode.json
     */
    /**
     * Sample code: Get a node.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getANode(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscNodes().getWithResponse("rg", "myAutomationAccount33", "nodeId", Context.NONE);
    }
}
