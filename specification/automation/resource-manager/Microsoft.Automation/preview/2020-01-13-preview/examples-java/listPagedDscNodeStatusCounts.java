import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.CountType;

/** Samples for NodeCountInformation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeStatusCounts.json
     */
    /**
     * Sample code: Get node's status counts.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getNodeSStatusCounts(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.nodeCountInformations().getWithResponse("rg", "myAutomationAccount33", CountType.STATUS, Context.NONE);
    }
}
