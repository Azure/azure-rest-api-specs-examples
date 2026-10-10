
/**
 * Samples for RecoveryResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryResources_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryResourcesListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryResources().list("sampleServiceGroupName", "plan1", com.azure.core.util.Context.NONE);
    }
}
