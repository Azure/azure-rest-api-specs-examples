
/**
 * Samples for ConnectionPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ConnectionPolicyList.json
     */
    /**
     * Sample code: ConnectionPolicyList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void connectionPolicyList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionPolicies().list("rg1", "TestHub", com.azure.core.util.Context.NONE);
    }
}
