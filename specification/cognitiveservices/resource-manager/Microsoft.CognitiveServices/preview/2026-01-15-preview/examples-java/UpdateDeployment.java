
import com.azure.resourcemanager.cognitiveservices.models.Deployment;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for Deployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/UpdateDeployment.json
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
