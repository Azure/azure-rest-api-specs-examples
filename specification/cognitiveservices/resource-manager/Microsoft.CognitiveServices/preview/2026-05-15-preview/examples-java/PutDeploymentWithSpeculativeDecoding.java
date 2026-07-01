
import com.azure.resourcemanager.cognitiveservices.models.DeploymentModel;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentProperties;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentSpeculativeDecoding;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentState;
import com.azure.resourcemanager.cognitiveservices.models.ServiceTier;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/PutDeploymentWithSpeculativeDecoding.json
     */
    /**
     * Sample code: PutDeploymentWithSpeculativeDecoding.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putDeploymentWithSpeculativeDecoding(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().define("deploymentName").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new DeploymentProperties()
                .withModel(new DeploymentModel().withFormat("Fireworks").withName("FW-Qwen3-14B").withVersion("1"))
                .withSpeculativeDecoding(new DeploymentSpeculativeDecoding().withDraftModel(new DeploymentModel()
                    .withFormat("FireworksCustom").withName("testDraftModel").withVersion("1").withSource(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/projects/projectName"))
                    .withDraftTokenCount(4))
                .withServiceTier(ServiceTier.DEFAULT).withDeploymentState(DeploymentState.RUNNING))
            .withSku(new Sku().withName("GlobalProvisionedManaged").withCapacity(80)).create();
    }
}
