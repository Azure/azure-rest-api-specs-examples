
/**
 * Samples for NeuroShieldFindings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldFindings_Get.json
     */
    /**
     * Sample code: Get finding.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void getFinding(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldFindings().getWithResponse("myResourceGroup", "myNeuroShield", "vm-01",
            "CVE-2026-12345-vm-01", com.azure.core.util.Context.NONE);
    }
}
