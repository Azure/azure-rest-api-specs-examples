
/**
 * Samples for ConnectionPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ConnectionPolicyDelete.json
     */
    /**
     * Sample code: ConnectionPolicyDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void connectionPolicyDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionPolicies().delete("rg1", "TestHub", "testpolicy",
            com.azure.core.util.Context.NONE);
    }
}
