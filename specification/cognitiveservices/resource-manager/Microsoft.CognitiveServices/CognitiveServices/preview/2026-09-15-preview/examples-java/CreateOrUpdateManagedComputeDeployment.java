
import com.azure.resourcemanager.cognitiveservices.models.DeploymentModelVersionUpgradeOption;
import com.azure.resourcemanager.cognitiveservices.models.GatedModelAccessProperties;
import com.azure.resourcemanager.cognitiveservices.models.ManagedComputeDeploymentProperties;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for ManagedComputeDeployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CreateOrUpdateManagedComputeDeployment.json
     */
    /**
     * Sample code: CreateOrUpdateManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createOrUpdateManagedComputeDeployment(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeDeployments().define("gpt-oss-120b-gpu")
            .withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new ManagedComputeDeploymentProperties()
                .withModel("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4")
                .withDeploymentTemplate(
                    "azureml://registries/azureml-openai-oss/deploymenttemplates/gpt-oss-120b-short-context/versions/1")
                .withAcceleratorType("H100_80GB")
                .withVersionUpgradeOption(DeploymentModelVersionUpgradeOption.ONCE_NEW_DEFAULT_VERSION_AVAILABLE)
                .withGatedModelAccess(new GatedModelAccessProperties().withConnectionId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/projects/my-project/connections/my-hf-connection")))
            .withSku(new Sku().withName("GlobalManagedCompute").withCapacity(1)).create();
    }
}
