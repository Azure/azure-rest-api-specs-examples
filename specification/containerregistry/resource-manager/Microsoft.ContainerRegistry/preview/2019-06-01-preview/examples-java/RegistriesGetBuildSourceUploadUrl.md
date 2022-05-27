Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Registries GetBuildSourceUploadUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesGetBuildSourceUploadUrl.json
     */
    /**
     * Sample code: Registries_GetBuildSourceUploadUrl.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registriesGetBuildSourceUploadUrl(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .getBuildSourceUploadUrlWithResponse("myResourceGroup", "myRegistry", Context.NONE);
    }
}
```
