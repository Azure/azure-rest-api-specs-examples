
/**
 * Samples for DdosCustomPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosCustomPolicyListAll.json
     */
    /**
     * Sample code: List all DDoS custom policies in subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllDDoSCustomPoliciesInSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosCustomPolicies().list(com.azure.core.util.Context.NONE);
    }
}
