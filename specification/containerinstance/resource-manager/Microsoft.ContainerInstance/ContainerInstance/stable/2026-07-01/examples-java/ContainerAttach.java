
/**
 * Samples for Containers Attach.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAttach.json
     */
    /**
     * Sample code: ContainerAttach.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerAttach(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainers().attachWithResponse("demo", "demo1", "container1",
            com.azure.core.util.Context.NONE);
    }
}
