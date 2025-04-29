
/**
 * Samples for Schedulers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/Schedulers_ListByResourceGroup.json
     */
    /**
     * Sample code: Schedulers_ListByResourceGroup.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersListByResourceGroup(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
