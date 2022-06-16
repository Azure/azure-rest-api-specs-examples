import com.azure.core.util.Context;

/** Samples for Quotas List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Quota/list.json
     */
    /**
     * Sample code: List workspace quotas by VMFamily.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listWorkspaceQuotasByVMFamily(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.quotas().list("eastus", Context.NONE);
    }
}
