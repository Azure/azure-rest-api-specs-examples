
/**
 * Samples for NGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsDelete.json
     */
    /**
     * Sample code: NGroupsDelete.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void nGroupsDelete(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().delete("demo", "demo-ngroup", com.azure.core.util.Context.NONE);
    }
}
