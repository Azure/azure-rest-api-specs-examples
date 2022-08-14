import com.azure.core.util.Context;

/** Samples for DscNode Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteDscNode.json
     */
    /**
     * Sample code: Delete a DSC Node.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteADSCNode(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .deleteWithResponse("rg", "myAutomationAccount9", "e1243a76-a9bd-432f-bde3-ad8f317ee786", Context.NONE);
    }
}
