```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.models.DeploymentMode;
import com.azure.resourcemanager.resources.models.DeploymentWhatIfProperties;
import com.azure.resourcemanager.resources.models.ScopedDeploymentWhatIf;
import com.azure.resourcemanager.resources.models.TemplateLink;
import java.io.IOException;

/** Samples for Deployments WhatIfAtManagementGroupScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PostDeploymentWhatIfOnManagementGroup.json
     */
    /**
     * Sample code: Predict template changes at management group scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void predictTemplateChangesAtManagementGroupScope(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getDeployments()
            .whatIfAtManagementGroupScope(
                "myManagementGruop",
                "exampleDeploymentName",
                new ScopedDeploymentWhatIf()
                    .withLocation("eastus")
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
