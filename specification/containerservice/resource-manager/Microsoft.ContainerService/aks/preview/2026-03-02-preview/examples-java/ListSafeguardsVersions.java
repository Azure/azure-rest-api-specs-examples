
/**
 * Samples for ManagedClusters ListSafeguardsVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ListSafeguardsVersions.json
     */
    /**
     * Sample code: List Safeguards Versions.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listSafeguardsVersions(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listSafeguardsVersions("location1",
            com.azure.core.util.Context.NONE);
    }
}
