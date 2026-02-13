
/**
 * Samples for Artifacts ListDownloadUri.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/Artifact_ListDownloadUri_MaximumSet_Gen.json
     */
    /**
     * Sample code: Artifacts_ListDownloadUri.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void artifactsListDownloadUri(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.artifacts().listDownloadUriWithResponse("rgdisconnectedoperations", "L4z_-S", "B-Ra--W0", "artifact1",
            com.azure.core.util.Context.NONE);
    }
}
