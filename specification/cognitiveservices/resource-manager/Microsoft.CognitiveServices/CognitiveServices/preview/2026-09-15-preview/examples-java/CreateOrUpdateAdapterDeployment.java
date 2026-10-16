
import com.azure.resourcemanager.cognitiveservices.models.AdapterDeploymentProperties;

/**
 * Samples for AdapterDeployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CreateOrUpdateAdapterDeployment.json
     */
    /**
     * Sample code: CreateOrUpdateAdapterDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        createOrUpdateAdapterDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.adapterDeployments().define("adapterDeploymentName")
            .withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new AdapterDeploymentProperties()
                .withSourceModelId("azureai://accounts/accountName/projects/projectName/models/modelName/versions/1")
                .withTargetDeploymentName("managedComputeDeploymentName"))
            .create();
    }
}
