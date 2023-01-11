import com.azure.core.util.Context;

/** Samples for Schedules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Schedule/list.json
     */
    /**
     * Sample code: List Schedules.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listSchedules(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.schedules().list("test-rg", "my-aml-workspace", null, null, Context.NONE);
    }
}
