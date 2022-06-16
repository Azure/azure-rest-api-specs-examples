import com.azure.resourcemanager.cognitiveservices.models.DeploymentModel;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentProperties;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentScaleSettings;
import com.azure.resourcemanager.cognitiveservices.models.DeploymentScaleType;

/** Samples for Deployments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/PutDeployment.json
     */
    /**
     * Sample code: PutDeployment.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .deployments()
            .define("deploymentName")
            .withExistingAccount("resourceGroupName", "accountName")
            .withProperties(
                new DeploymentProperties()
                    .withModel(new DeploymentModel().withFormat("OpenAI").withName("ada").withVersion("1"))
                    .withScaleSettings(
                        new DeploymentScaleSettings().withScaleType(DeploymentScaleType.MANUAL).withCapacity(1)))
            .create();
    }
}
