
/**
 * Samples for ManagedClusters ListNodeImageVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/NodeImageVersions_List.json
     */
    /**
     * Sample code: List Node Image Versions.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listNodeImageVersions(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listNodeImageVersions("location1",
            com.azure.core.util.Context.NONE);
    }
}
