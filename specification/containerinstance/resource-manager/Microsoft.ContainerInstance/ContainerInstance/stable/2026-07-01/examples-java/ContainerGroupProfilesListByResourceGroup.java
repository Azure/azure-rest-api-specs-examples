
/**
 * Samples for CGProfilesOperation ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfilesListByResourceGroup.json
     */
    /**
     * Sample code: ContainerGroupProfilesListByResourceGroup.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfilesListByResourceGroup(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfilesOperations().listByResourceGroup("demo", com.azure.core.util.Context.NONE);
    }
}
