
/**
 * Samples for Schedulers ListPrivateLinks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PrivateLinkResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateLinkResources_List_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        privateLinkResourcesListMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().listPrivateLinks("rgdurabletask", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
