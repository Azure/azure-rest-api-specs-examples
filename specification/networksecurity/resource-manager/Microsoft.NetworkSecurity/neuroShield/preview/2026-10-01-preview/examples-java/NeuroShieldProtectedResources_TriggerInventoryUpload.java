
/**
 * Samples for NeuroShieldProtectedResources TriggerInventoryUpload.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldProtectedResources_TriggerInventoryUpload.json
     */
    /**
     * Sample code: Trigger an inventory upload.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        triggerAnInventoryUpload(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldProtectedResources().triggerInventoryUploadWithResponse("myResourceGroup", "myNeuroShield",
            "vm-01", com.azure.core.util.Context.NONE);
    }
}
