```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.CustomDomainResourceInner;
import com.azure.resourcemanager.appplatform.models.CustomDomainProperties;

/** Samples for CustomDomains CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/CustomDomains_CreateOrUpdate.json
     */
    /**
     * Sample code: CustomDomains_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .createOrUpdate(
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
