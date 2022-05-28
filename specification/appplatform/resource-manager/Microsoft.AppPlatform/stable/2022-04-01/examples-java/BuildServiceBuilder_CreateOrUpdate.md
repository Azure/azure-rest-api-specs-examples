Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.BuilderResourceInner;
import com.azure.resourcemanager.appplatform.models.BuilderProperties;
import com.azure.resourcemanager.appplatform.models.BuildpackProperties;
import com.azure.resourcemanager.appplatform.models.BuildpacksGroupProperties;
import com.azure.resourcemanager.appplatform.models.StackProperties;
import java.util.Arrays;

/** Samples for BuildServiceBuilder CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildServiceBuilder_CreateOrUpdate.json
     */
    /**
     * Sample code: BuildServiceBuilder_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceBuilderCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServiceBuilders()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "default",
                "mybuilder",
                new BuilderResourceInner()
                    .withProperties(
                        new BuilderProperties()
                            .withStack(new StackProperties().withId("io.buildpacks.stacks.bionic").withVersion("base"))
                            .withBuildpackGroups(
                                Arrays
                                    .asList(
                                        new BuildpacksGroupProperties()
                                            .withName("mix")
                                            .withBuildpacks(
                                                Arrays
                                                    .asList(
                                                        new BuildpackProperties()
                                                            .withId("tanzu-buildpacks/java-azure")))))),
                Context.NONE);
    }
}
```
