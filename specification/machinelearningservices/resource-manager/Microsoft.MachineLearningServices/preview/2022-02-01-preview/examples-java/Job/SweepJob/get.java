import com.azure.core.util.Context;

/** Samples for Jobs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/SweepJob/get.json
     */
    /**
     * Sample code: Get Sweep Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getSweepJob(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.jobs().getWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
