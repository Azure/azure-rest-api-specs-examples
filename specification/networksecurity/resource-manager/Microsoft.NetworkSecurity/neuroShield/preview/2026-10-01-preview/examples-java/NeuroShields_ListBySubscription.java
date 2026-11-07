
/**
 * Samples for NeuroShields List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShields_ListBySubscription.json
     */
    /**
     * Sample code: List NeuroShield resources by subscription.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void listNeuroShieldResourcesBySubscription(
        com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShields().list(null, null, com.azure.core.util.Context.NONE);
    }
}
