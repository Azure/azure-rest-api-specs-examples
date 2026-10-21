
/**
 * Samples for Schedulers Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Schedulers_Restart.json
     */
    /**
     * Sample code: Schedulers_Restart.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersRestart(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().restartWithResponse("rgopenapi", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
