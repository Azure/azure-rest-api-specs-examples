
/**
 * Samples for NeuroShieldFindings Rollback.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldFindings_Rollback.json
     */
    /**
     * Sample code: Rollback finding mitigations.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        rollbackFindingMitigations(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldFindings().rollbackWithResponse("myResourceGroup", "myNeuroShield", "vm-01",
            "CVE-2026-12345-vm-01", com.azure.core.util.Context.NONE);
    }
}
