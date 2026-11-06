
/**
 * Samples for NeuroShieldFindings Mitigate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldFindings_Mitigate.json
     */
    /**
     * Sample code: Mitigate finding.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void mitigateFinding(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldFindings().mitigateWithResponse("myResourceGroup", "myNeuroShield", "vm-01",
            "CVE-2026-12345-vm-01", com.azure.core.util.Context.NONE);
    }
}
