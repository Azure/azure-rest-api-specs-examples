
/**
 * Samples for CGProfile GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfilesGet.json
     */
    /**
     * Sample code: ContainerGroupProfilesGet_Succeeded.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfilesGetSucceeded(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfiles().getByResourceGroupWithResponse("demo", "demo1",
            com.azure.core.util.Context.NONE);
    }
}
