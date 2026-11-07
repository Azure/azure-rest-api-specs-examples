
/**
 * Samples for NeuroShieldMitigations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_Get_Queued.json
     */
    /**
     * Sample code: Get mitigation (Queued status).
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        getMitigationQueuedStatus(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldMitigations().getWithResponse("myResourceGroup", "myNeuroShield", "vm-01", "mitigation-01",
            com.azure.core.util.Context.NONE);
    }
}
