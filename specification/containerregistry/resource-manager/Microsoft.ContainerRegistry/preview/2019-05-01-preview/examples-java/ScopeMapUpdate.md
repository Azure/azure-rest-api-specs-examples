```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.ScopeMapUpdateParameters;
import java.util.Arrays;

/** Samples for ScopeMaps Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/ScopeMapUpdate.json
     */
    /**
     * Sample code: ScopeMapUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getScopeMaps()
            .update(
                "myResourceGroup",
                "myRegistry",
                "myScopeMap",
                new ScopeMapUpdateParameters()
                    .withDescription("Developer Scopes")
                    .withActions(
                        Arrays
                            .asList("repositories/myrepository/contentWrite", "repositories/myrepository/contentRead")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
