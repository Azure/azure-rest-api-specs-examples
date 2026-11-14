
/**
 * Samples for ContainerGroups Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsStart.json
     */
    /**
     * Sample code: ContainerStart.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerStart(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().start("demo", "demo1", com.azure.core.util.Context.NONE);
    }
}
