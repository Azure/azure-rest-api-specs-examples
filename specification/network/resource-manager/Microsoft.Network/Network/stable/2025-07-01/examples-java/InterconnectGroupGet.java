
/**
 * Samples for InterconnectGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InterconnectGroupGet.json
     */
    /**
     * Sample code: Get interconnect group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getInterconnectGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInterconnectGroups().getByResourceGroupWithResponse("rg1", "test-ig",
            com.azure.core.util.Context.NONE);
    }
}
