
/**
 * Samples for DisconnectedOperations ListDeploymentManifest.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DisconnectedOperations_ListDeploymentManifest_MaximumSet_Gen.json
     */
    /**
     * Sample code: DisconnectedOperations_ListDeploymentManifest.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void disconnectedOperationsListDeploymentManifest(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.disconnectedOperations().listDeploymentManifestWithResponse("rgdisconnectedoperations", "demo-resource",
            com.azure.core.util.Context.NONE);
    }
}
