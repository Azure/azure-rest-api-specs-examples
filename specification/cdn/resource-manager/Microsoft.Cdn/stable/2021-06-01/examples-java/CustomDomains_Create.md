Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.CustomDomainParameters;

/** Samples for CustomDomains Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_Create.json
     */
    /**
     * Sample code: CustomDomains_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .create(
                "RG",
                "profile1",
                "endpoint1",
                "www-someDomain-net",
                new CustomDomainParameters().withHostname("www.someDomain.net"),
                Context.NONE);
    }
}
```
