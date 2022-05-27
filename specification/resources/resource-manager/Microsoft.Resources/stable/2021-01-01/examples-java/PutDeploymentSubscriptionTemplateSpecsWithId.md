Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.DeploymentInner;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentProperties;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.io.IOException;

/** Samples for Deployments CreateOrUpdateAtSubscriptionScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PutDeploymentSubscriptionTemplateSpecsWithId.json
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
            .createOrUpdateAtSubscriptionScope(
                "my-deployment",
                new DeploymentInner()
                    .withLocation("eastus")
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
                Context.NONE);
    }
}
```
