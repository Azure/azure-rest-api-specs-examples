
/**
 * Samples for Connector DeleteDryrun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ConnectorDryrunDelete.json
     */
    /**
     * Sample code: ConnectorDryrunDelete.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void connectorDryrunDelete(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().deleteDryrunWithResponse("00000000-0000-0000-0000-000000000000", "test-rg", "westus",
            "dryrunName", com.azure.core.util.Context.NONE);
    }
}
