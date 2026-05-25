
/**
 * Samples for ManagedClusters GetSafeguardsVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/GetSafeguardsVersions.json
     */
    /**
     * Sample code: Get Safeguards available versions.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getSafeguardsAvailableVersions(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().getSafeguardsVersionsWithResponse("location1", "v1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
