
/**
 * Samples for ContainerGroups Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsStop.json
     */
    /**
     * Sample code: ContainerStop.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerStop(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().stopWithResponse("demo", "demo1",
            com.azure.core.util.Context.NONE);
    }
}
