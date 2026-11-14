
/**
 * Samples for CGProfile Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfilesDelete.json
     */
    /**
     * Sample code: ContainerGroupProfilesDelete.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        containerGroupProfilesDelete(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfiles().deleteWithResponse("demo", "demo1", com.azure.core.util.Context.NONE);
    }
}
