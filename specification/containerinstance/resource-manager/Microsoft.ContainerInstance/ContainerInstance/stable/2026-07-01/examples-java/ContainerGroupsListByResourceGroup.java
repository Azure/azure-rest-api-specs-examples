
/**
 * Samples for ContainerGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsListByResourceGroup.json
     */
    /**
     * Sample code: ContainerGroupsListByResourceGroup.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupsListByResourceGroup(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().listByResourceGroup("demo", com.azure.core.util.Context.NONE);
    }
}
