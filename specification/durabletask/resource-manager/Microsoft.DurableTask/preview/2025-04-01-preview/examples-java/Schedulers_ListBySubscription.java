
/**
 * Samples for Schedulers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/Schedulers_ListBySubscription.json
     */
    /**
     * Sample code: Schedulers_ListBySubscription.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersListBySubscription(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().list(com.azure.core.util.Context.NONE);
    }
}
