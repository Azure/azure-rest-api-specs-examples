import com.azure.core.util.Context;

/** Samples for Deployments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/GetDeployment.json
     */
    /**
     * Sample code: GetDeployment.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().getWithResponse("resourceGroupName", "accountName", "deploymentName", Context.NONE);
    }
}
