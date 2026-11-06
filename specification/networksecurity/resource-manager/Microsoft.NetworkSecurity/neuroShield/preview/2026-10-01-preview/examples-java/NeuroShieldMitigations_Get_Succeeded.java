
/**
 * Samples for NeuroShieldMitigations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_Get_Succeeded.json
     */
    /**
     * Sample code: Get mitigation (Succeeded status).
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        getMitigationSucceededStatus(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldMitigations().getWithResponse("myResourceGroup", "myNeuroShield", "vm-01", "mitigation-01",
            com.azure.core.util.Context.NONE);
    }
}
