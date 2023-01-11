import com.azure.core.util.Context;

/** Samples for Schedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Schedule/delete.json
     */
    /**
     * Sample code: Delete Schedule.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteSchedule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.schedules().delete("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
