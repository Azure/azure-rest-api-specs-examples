import com.azure.core.util.Context;

/** Samples for Jobs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/PipelineJob/get.json
     */
    /**
     * Sample code: Get Pipeline Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getPipelineJob(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.jobs().getWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
