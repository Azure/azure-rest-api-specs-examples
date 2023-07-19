/** Samples for Deployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListDeployments.json
     */
    /**
     * Sample code: ListDeployments.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
