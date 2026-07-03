
/**
 * Samples for ConnectionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ConnectionPolicyGet.json
     */
    /**
     * Sample code: ConnectionPolicyGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void connectionPolicyGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConnectionPolicies().getWithResponse("rg1", "TestHub", "testpolicy",
            com.azure.core.util.Context.NONE);
    }
}
