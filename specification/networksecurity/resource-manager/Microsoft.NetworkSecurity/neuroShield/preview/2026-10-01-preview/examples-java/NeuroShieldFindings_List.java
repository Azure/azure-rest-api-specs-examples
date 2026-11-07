
/**
 * Samples for NeuroShieldFindings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldFindings_List.json
     */
    /**
     * Sample code: List findings.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void listFindings(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldFindings().list("myResourceGroup", "myNeuroShield", "vm-01", null, null,
            com.azure.core.util.Context.NONE);
    }
}
