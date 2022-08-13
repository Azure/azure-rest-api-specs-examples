import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.CountType;

/** Samples for NodeCountInformation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeConfigurationCounts.json
     */
    /**
     * Sample code: Get node's node configuration counts.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getNodeSNodeConfigurationCounts(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .nodeCountInformations()
            .getWithResponse("rg", "myAutomationAccount33", CountType.NODECONFIGURATION, Context.NONE);
    }
}
