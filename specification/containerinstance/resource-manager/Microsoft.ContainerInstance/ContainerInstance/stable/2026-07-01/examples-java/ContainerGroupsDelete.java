
/**
 * Samples for ContainerGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsDelete.json
     */
    /**
     * Sample code: ContainerGroupsDelete.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        containerGroupsDelete(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().delete("demo", "demo1", com.azure.core.util.Context.NONE);
    }
}
