
/**
 * Samples for Schedules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Schedule/delete.json
     */
    /**
     * Sample code: Delete Schedule.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteSchedule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.schedules().delete("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
