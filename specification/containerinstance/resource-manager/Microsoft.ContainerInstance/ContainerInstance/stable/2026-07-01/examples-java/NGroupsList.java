
/**
 * Samples for NGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/NGroupsList.json
     */
    /**
     * Sample code: NGroupsList.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void nGroupsList(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getNGroups().list(com.azure.core.util.Context.NONE);
    }
}
