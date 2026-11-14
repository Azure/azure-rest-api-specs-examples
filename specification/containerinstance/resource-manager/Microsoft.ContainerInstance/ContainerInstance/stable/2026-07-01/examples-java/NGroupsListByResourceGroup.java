
/**
 * Samples for NGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsListByResourceGroup.json
     */
    /**
     * Sample code: NGroupsListByResourceGroup.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        nGroupsListByResourceGroup(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().listByResourceGroup("demo", com.azure.core.util.Context.NONE);
    }
}
