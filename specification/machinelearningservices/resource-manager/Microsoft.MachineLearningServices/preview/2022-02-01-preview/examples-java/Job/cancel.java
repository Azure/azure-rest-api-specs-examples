import com.azure.core.util.Context;

/** Samples for Jobs Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/cancel.json
     */
    /**
     * Sample code: Cancel Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void cancelJob(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.jobs().cancelWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
