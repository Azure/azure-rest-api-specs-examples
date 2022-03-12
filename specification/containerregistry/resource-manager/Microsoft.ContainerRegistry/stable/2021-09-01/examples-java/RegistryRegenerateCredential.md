Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.PasswordName;
import com.azure.resourcemanager.containerregistry.models.RegenerateCredentialParameters;

/** Samples for Registries RegenerateCredential. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/RegistryRegenerateCredential.json
     */
    /**
     * Sample code: RegistryRegenerateCredential.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryRegenerateCredential(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .regenerateCredentialWithResponse(
                "myResourceGroup",
                "myRegistry",
                new RegenerateCredentialParameters().withName(PasswordName.PASSWORD),
                Context.NONE);
    }
}
```
