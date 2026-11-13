
/**
 * Samples for NGroups Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsStart.json
     */
    /**
     * Sample code: NGroupsStart.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void nGroupsStart(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().start("demo", "demo-ngroup", com.azure.core.util.Context.NONE);
    }
}
