Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.CustomDomainResourceInner;
import com.azure.resourcemanager.appplatform.models.CustomDomainProperties;

/** Samples for CustomDomains Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/CustomDomains_Update.json
     */
    /**
     * Sample code: CustomDomains_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .update(
                "myResourceGroup",
                "myservice",
                "myapp",
                "mydomain.com",
                new CustomDomainResourceInner()
                    .withProperties(
                        new CustomDomainProperties()
                            .withThumbprint("934367bf1c97033f877db0f15cb1b586957d3133")
                            .withCertName("mycert")),
                Context.NONE);
    }
}
```
