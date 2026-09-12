
/**
 * Samples for HybridConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionGet.json
     */
    /**
     * Sample code: RelayHybridConnectionGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().getWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", com.azure.core.util.Context.NONE);
    }
}
