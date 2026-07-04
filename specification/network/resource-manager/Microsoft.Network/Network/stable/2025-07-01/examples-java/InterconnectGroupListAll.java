
/**
 * Samples for InterconnectGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InterconnectGroupListAll.json
     */
    /**
     * Sample code: List all interconnect groups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllInterconnectGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInterconnectGroups().list(com.azure.core.util.Context.NONE);
    }
}
