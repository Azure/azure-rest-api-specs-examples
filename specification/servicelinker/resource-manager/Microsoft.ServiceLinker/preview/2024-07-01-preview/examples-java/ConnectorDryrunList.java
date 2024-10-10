
/**
 * Samples for Connector ListDryrun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ConnectorDryrunList.json
     */
    /**
     * Sample code: ConnectorDryrunList.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void connectorDryrunList(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().listDryrun("00000000-0000-0000-0000-000000000000", "test-rg", "westus",
            com.azure.core.util.Context.NONE);
    }
}
