
/**
 * Samples for Connector Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Connector_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Connector_Delete_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void connectorDeleteMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.connectors().delete("rgconfluent", "frwocpndztguhgng", "duq", "chw", "suaugvwtvhexoqdrmxknvyiobq",
            com.azure.core.util.Context.NONE);
    }
}
