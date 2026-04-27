
/**
 * Samples for Connector Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Connector_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Connector_Delete_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void connectorDeleteMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.connectors().delete("rgconfluent", "xqspbodq", "aabxehocioujmjjkgegijsmntw", "seivpzvrbyhjfmqb",
            "qznabwwh", com.azure.core.util.Context.NONE);
    }
}
