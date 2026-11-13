
import com.azure.resourcemanager.cognitiveservices.models.DeploymentModel;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentProperties;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentState;
import com.azure.resourcemanager.cognitiveservices.models.ServiceTier;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/PutDeploymentWithSpeculativeDecoding.json
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
                .withServiceTier(ServiceTier.DEFAULT).withDeploymentState(DeploymentState.RUNNING))
            .withSku(new Sku().withName("GlobalProvisionedManaged").withCapacity(80)).create();
    }
}
