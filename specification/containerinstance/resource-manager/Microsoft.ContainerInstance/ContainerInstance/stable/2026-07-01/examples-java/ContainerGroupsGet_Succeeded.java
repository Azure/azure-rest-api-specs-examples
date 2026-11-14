
/**
 * Samples for ContainerGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupsGet_Succeeded.json
     */
    /**
     * Sample code: ContainerGroupsGet_Succeeded.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        containerGroupsGetSucceeded(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainerGroups().getByResourceGroupWithResponse("demo", "demo1",
            com.azure.core.util.Context.NONE);
    }
}
