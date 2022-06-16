import com.azure.core.util.Context;

/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/SweepJob/list.json
     */
    /**
     * Sample code: List Sweep Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listSweepJob(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.jobs().list("test-rg", "my-aml-workspace", null, "string", "string", null, null, null, Context.NONE);
    }
}
