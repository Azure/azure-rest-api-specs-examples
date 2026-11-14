
/**
 * Samples for NGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsGet.json
     */
    /**
     * Sample code: NGroupsGet.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void nGroupsGet(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().getByResourceGroupWithResponse("demo", "demo-ngroup",
            com.azure.core.util.Context.NONE);
    }
}
