import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.DeploymentInner;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import com.azure.resourcemanager.resources.models.OnErrorDeployment;
import com.azure.resourcemanager.resources.models.OnErrorDeploymentType;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.io.IOException;

/** Samples for Deployments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PutDeploymentWithOnErrorDeploymentSpecificDeployment.json
     */
    /**
     * Sample code: Create a deployment that will redeploy another deployment on failure.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADeploymentThatWillRedeployAnotherDeploymentOnFailure(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getDeployments()
            .createOrUpdate(
                "my-resource-group",
                "my-deployment",
                new DeploymentInner()
                    .withProperties(
                        new DeploymentProperties()
                            .withTemplateLink(new TemplateLink().withUri("https://example.com/exampleTemplate.json"))
                            .withParameters(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
                            .withMode(DeploymentMode.COMPLETE)
                            .withOnErrorDeployment(
                                new OnErrorDeployment()
                                    .withType(OnErrorDeploymentType.SPECIFIC_DEPLOYMENT)
                                    .withDeploymentName("name-of-deployment-to-use"))),
                com.azure.core.util.Context.NONE);
    }
}
