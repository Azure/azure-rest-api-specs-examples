
/**
 * Samples for CGProfile ListAllRevisions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfileListAllRevisions.json
     */
    /**
     * Sample code: ContainerGroupProfileListAllRevisions.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfileListAllRevisions(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfiles().listAllRevisions("demo", "demo1", com.azure.core.util.Context.NONE);
    }
}
