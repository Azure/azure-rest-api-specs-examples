Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentWhatIf;
import com.azure.resourcemanager.resources.models.DeploymentWhatIfProperties;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.io.IOException;

/** Samples for Deployments WhatIf. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PostDeploymentWhatIfOnResourceGroup.json
     */
    /**
     * Sample code: Predict template changes at resource group scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void predictTemplateChangesAtResourceGroupScope(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getDeployments()
            .whatIf(
                "my-resource-group",
                "my-deployment",
                new DeploymentWhatIf()
                    .withProperties(
                        new DeploymentWhatIfProperties()
                            .withTemplateLink(new TemplateLink())
                            .withParameters(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
                            .withMode(DeploymentMode.INCREMENTAL)),
                Context.NONE);
    }
}
```
