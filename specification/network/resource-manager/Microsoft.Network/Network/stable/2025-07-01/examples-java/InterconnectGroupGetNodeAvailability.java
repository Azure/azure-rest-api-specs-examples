
/**
 * Samples for InterconnectGroups GetNodeAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InterconnectGroupGetNodeAvailability.json
     */
    /**
     * Sample code: Get interconnect group node availability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getInterconnectGroupNodeAvailability(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInterconnectGroups().getNodeAvailability("rg1", "test-ig",
            com.azure.core.util.Context.NONE);
    }
}
