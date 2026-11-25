
/**
 * Samples for Quotas List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Quota/list.json
     */
    /**
     * Sample code: List workspace quotas by VMFamily.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceQuotasByVMFamily(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.quotas().list("eastus", com.azure.core.util.Context.NONE);
    }
}
