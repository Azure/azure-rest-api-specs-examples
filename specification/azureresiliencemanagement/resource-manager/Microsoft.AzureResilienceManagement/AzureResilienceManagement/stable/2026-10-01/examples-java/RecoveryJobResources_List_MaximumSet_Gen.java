
/**
 * Samples for RecoveryJobResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryJobResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryJobResources_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryJobResourcesListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryJobResources().list("sampleServiceGroupName", "samplePlanName",
            "c56888ef-9ced-4001-a6d4-7145a0309bdb", com.azure.core.util.Context.NONE);
    }
}
