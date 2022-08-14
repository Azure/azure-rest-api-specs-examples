import com.azure.core.util.Context;

/** Samples for DscNodeConfiguration Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteDscNodeConfiguration.json
     */
    /**
     * Sample code: Delete a DSC node configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteADSCNodeConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodeConfigurations()
            .deleteWithResponse("rg", "myAutomationAccount20", "configName.nodeConfigName", Context.NONE);
    }
}
