
/**
 * Samples for ProjectConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ProjectConnection/delete.json
     */
    /**
     * Sample code: DeleteProjectConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteProjectConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectConnections().deleteWithResponse("resourceGroup-1", "account-1", "project-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}
