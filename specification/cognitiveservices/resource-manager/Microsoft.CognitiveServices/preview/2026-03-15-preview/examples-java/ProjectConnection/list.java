
/**
 * Samples for ProjectConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ProjectConnection/list.json
     */
    /**
     * Sample code: ListProjectConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listProjectConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectConnections().list("resourceGroup-1", "account-1", "project-1", "[target url]",
            "ContainerRegistry", null, com.azure.core.util.Context.NONE);
    }
}
