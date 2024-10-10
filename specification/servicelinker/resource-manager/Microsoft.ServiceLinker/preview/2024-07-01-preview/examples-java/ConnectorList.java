
/**
 * Samples for Connector List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ConnectorList.json
     */
    /**
     * Sample code: ConnectorList.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void connectorList(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().list("00000000-0000-0000-0000-000000000000", "test-rg", "westus",
            com.azure.core.util.Context.NONE);
    }
}
