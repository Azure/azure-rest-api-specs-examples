import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.DscNodeUpdateParameters;
import com.azure.resourcemanager.automation.models.DscNodeUpdateParametersProperties;

/** Samples for DscNode Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateDscNode.json
     */
    /**
     * Sample code: Update a node.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateANode(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .updateWithResponse(
                "rg",
                "myAutomationAccount33",
                "nodeId",
                new DscNodeUpdateParameters()
                    .withNodeId("nodeId")
                    .withProperties(new DscNodeUpdateParametersProperties().withName("SetupServer.localhost")),
                Context.NONE);
    }
}
