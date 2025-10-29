
/**
 * Samples for ProjectConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ProjectConnection/get.json
     */
    /**
     * Sample code: GetProjectConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getProjectConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectConnections().getWithResponse("resourceGroup-1", "account-1", "project-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}
