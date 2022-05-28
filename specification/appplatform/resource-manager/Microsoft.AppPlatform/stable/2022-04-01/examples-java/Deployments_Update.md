Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.DeploymentResourceInner;
import com.azure.resourcemanager.appplatform.models.DeploymentResourceProperties;
import com.azure.resourcemanager.appplatform.models.SourceUploadedUserSourceInfo;

/** Samples for Deployments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_Update.json
     */
    /**
     * Sample code: Deployments_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .update(
                "myResourceGroup",
                "myservice",
                "myapp",
                "mydeployment",
                new DeploymentResourceInner()
                    .withProperties(
                        new DeploymentResourceProperties()
                            .withSource(
                                new SourceUploadedUserSourceInfo()
                                    .withVersion("1.0")
                                    .withRelativePath(
                                        "resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc")
                                    .withArtifactSelector("sub-module-1"))),
                Context.NONE);
    }
}
```
