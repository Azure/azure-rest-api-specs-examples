
/**
 * Samples for NeuroShieldMitigations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_List.json
     */
    /**
     * Sample code: List mitigations.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void listMitigations(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldMitigations().list("myResourceGroup", "myNeuroShield", "vm-01", null, null,
            com.azure.core.util.Context.NONE);
    }
}
