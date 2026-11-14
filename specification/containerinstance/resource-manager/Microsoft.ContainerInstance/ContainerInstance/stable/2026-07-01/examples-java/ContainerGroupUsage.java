
/**
 * Samples for Location ListUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupUsage.json
     */
    /**
     * Sample code: ContainerUsage.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerUsage(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getLocations().listUsage("westcentralus", com.azure.core.util.Context.NONE);
    }
}
