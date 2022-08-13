import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("rg1", "ddb1", "privateEndpointConnectionName", Context.NONE);
    }
}
