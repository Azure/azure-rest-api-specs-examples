
/**
 * Samples for NGroups Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsStop.json
     */
    /**
     * Sample code: NGroupsStop.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void nGroupsStop(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().stopWithResponse("demo", "demo-ngroup", com.azure.core.util.Context.NONE);
    }
}
