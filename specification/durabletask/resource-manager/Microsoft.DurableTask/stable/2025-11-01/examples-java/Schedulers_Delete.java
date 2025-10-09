
/**
 * Samples for Schedulers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/Schedulers_Delete.json
     */
    /**
     * Sample code: Schedulers_Delete.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersDelete(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().delete("rgopenapi", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
