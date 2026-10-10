
/**
 * Samples for RecoveryResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryResources_Get_CrossZoneVMRecovery.json
     */
    /**
     * Sample code: RecoveryResources_Get_CrossZoneVMRecovery.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryResourcesGetCrossZoneVMRecovery(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryResources().getWithResponse("sampleServiceGroupName", "samplePlanName",
            "12345678-9012-3456-7890-123456789012", com.azure.core.util.Context.NONE);
    }
}
