
import com.azure.resourcemanager.cognitiveservices.models.ManagedComputeDeploymentProperties;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for ManagedComputeDeployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/CreateOrUpdateVmManagedComputeDeployment.json
     */
    /**
     * Sample code: CreateOrUpdateVmManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createOrUpdateVmManagedComputeDeployment(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeDeployments().define("gpt-oss-120b-byoc")
            .withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new ManagedComputeDeploymentProperties()
                .withModel("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4")
                .withDeploymentTemplate("projects/my-project/deploymentTemplates/gpt-oss-120b-vllm-tuned/versions/2")
                .withComputeId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/computes/my-h100-pool")
                .withPriority("High"))
            .withSku(new Sku().withName("VmManagedCompute").withCapacity(2)).create();
    }
}
