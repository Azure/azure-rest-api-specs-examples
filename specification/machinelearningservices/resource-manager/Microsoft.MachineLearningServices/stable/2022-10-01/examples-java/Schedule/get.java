import com.azure.core.util.Context;

/** Samples for Schedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Schedule/get.json
     */
    /**
     * Sample code: Get Schedule.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getSchedule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.schedules().getWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
