
/**
 * Samples for Detectors ListByScheduler.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Detectors_ListByScheduler.json
     */
    /**
     * Sample code: Detectors_ListByScheduler.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void detectorsListByScheduler(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.detectors().listByScheduler("rgopenapi", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
