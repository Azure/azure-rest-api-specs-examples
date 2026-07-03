
/**
 * Samples for InterconnectGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InterconnectGroupDelete.json
     */
    /**
     * Sample code: Delete interconnect group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteInterconnectGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInterconnectGroups().deleteWithResponse("rg1", "test-ig",
            com.azure.core.util.Context.NONE);
    }
}
