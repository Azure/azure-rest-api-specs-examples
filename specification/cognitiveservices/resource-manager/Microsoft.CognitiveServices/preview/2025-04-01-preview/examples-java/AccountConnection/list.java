
/**
 * Samples for AccountConnection List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * AccountConnection/list.json
     */
    /**
     * Sample code: ListAccountConnections.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountConnections(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountConnections().list("resourceGroup-1", "account-1", "[tartget url]", "ContainerRegistry", null,
            com.azure.core.util.Context.NONE);
    }
}
