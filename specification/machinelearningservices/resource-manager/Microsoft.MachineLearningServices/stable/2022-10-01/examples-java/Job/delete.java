/** Samples for Jobs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/delete.json
     */
    /**
     * Sample code: Delete Job.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().delete("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
