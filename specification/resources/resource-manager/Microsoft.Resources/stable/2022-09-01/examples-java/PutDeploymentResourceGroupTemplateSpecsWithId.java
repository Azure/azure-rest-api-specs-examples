import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.DeploymentInner;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.io.IOException;

/** Samples for Deployments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/PutDeploymentResourceGroupTemplateSpecsWithId.json
     */
    /**
     * Sample code: Create a deployment that will deploy a templateSpec with the given resourceId.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADeploymentThatWillDeployATemplateSpecWithTheGivenResourceId(
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
                            .withTemplateLink(
                                new TemplateLink()
                                    .withId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1"))
                            .withParameters(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
                            .withMode(DeploymentMode.INCREMENTAL)),
                com.azure.core.util.Context.NONE);
    }
}
