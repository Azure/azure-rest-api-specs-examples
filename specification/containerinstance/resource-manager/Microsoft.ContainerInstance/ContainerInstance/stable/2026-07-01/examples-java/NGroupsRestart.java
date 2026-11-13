
/**
 * Samples for NGroups Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsRestart.json
     */
    /**
     * Sample code: NGroupsRestart.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void nGroupsRestart(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().restart("demo", "demo-ngroup", com.azure.core.util.Context.NONE);
    }
}
