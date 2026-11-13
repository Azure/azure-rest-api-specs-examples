
/**
 * Samples for ContainerGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsList.json
     */
    /**
     * Sample code: ContainerGroupsList.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        containerGroupsList(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().list(com.azure.core.util.Context.NONE);
    }
}
