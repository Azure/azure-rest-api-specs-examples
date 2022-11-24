import com.azure.core.util.Context;

/** Samples for Deployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/ListDeployments.json
     */
    /**
     * Sample code: ListDeployments.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().list("resourceGroupName", "accountName", Context.NONE);
    }
}
