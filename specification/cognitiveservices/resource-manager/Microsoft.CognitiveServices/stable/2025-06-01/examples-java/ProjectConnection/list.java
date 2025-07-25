
/**
 * Samples for ProjectConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ProjectConnection/list.json
     */
    /**
     * Sample code: ListProjectConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listProjectConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectConnections().list("resourceGroup-1", "account-1", "project-1", "[tartget url]",
            "ContainerRegistry", null, com.azure.core.util.Context.NONE);
    }
}
