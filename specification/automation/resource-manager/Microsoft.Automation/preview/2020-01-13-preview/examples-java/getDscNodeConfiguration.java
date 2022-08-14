import com.azure.core.util.Context;

/** Samples for DscNodeConfiguration Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeConfiguration.json
     */
    /**
     * Sample code: Get a DSC node configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getADSCNodeConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodeConfigurations()
            .getWithResponse("rg", "myAutomationAccount33", "SetupServer.localhost", Context.NONE);
    }
}
