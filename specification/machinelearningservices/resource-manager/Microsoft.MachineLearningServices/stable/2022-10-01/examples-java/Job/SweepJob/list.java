/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/SweepJob/list.json
     */
    /**
     * Sample code: List Sweep Job.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listSweepJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .jobs()
            .list("test-rg", "my-aml-workspace", null, "string", "string", null, com.azure.core.util.Context.NONE);
    }
}
