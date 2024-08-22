
/**
 * Samples for Quotas List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Quota/list.json
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
