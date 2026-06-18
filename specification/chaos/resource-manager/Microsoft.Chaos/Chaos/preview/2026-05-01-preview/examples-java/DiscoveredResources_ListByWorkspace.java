
/**
 * Samples for DiscoveredResources ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/DiscoveredResources_ListByWorkspace.json
     */
    /**
     * Sample code: Get a list of discovered resources for a workspace.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getAListOfDiscoveredResourcesForAWorkspace(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.discoveredResources().listByWorkspace("exampleRG", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
