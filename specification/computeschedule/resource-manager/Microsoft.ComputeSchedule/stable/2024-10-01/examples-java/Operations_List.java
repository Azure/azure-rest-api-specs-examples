
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void operationsList(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
