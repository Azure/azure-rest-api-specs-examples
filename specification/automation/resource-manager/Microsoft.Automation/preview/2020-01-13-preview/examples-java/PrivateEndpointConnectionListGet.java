import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateEndpointConnectionListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.privateEndpointConnections().listByAutomationAccount("rg1", "ddb1", Context.NONE);
    }
}
