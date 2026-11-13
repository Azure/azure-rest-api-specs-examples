
/**
 * Samples for CGProfile GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfilesGetPriority.json
     */
    /**
     * Sample code: ContainerGroupProfilesGetWithPriority.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfilesGetWithPriority(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfiles().getByResourceGroupWithResponse("demo", "demo1",
            com.azure.core.util.Context.NONE);
    }
}
