
import com.azure.resourcemanager.cognitiveservices.models.Deployment;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for Deployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * UpdateDeployment.json
     */
    /**
     * Sample code: UpdateDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Deployment resource = manager.deployments()
            .getWithResponse("resourceGroupName", "accountName", "deploymentName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withSku(new Sku().withName("Standard").withCapacity(1)).apply();
    }
}
