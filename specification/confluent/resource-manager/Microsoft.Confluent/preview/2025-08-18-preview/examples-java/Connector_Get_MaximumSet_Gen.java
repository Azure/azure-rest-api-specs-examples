
/**
 * Samples for Connector Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Connector_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Connector_Get_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void connectorGetMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.connectors().getWithResponse("rgconfluent", "pgwuoi", "rxbrvvdnplvbedrzwbgtwhbdm",
            "eknmpvbhtvwxdxddkos", "zakwjragxeiur", com.azure.core.util.Context.NONE);
    }
}
