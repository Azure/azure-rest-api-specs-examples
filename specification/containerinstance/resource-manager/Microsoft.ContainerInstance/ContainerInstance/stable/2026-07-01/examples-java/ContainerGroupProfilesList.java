
/**
 * Samples for CGProfilesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfilesList.json
     */
    /**
     * Sample code: ContainerGroupProfilesList.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void
        containerGroupProfilesList(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfilesOperations().list(com.azure.core.util.Context.NONE);
    }
}
