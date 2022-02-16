Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.ValidateCustomDomainInput;

/** Samples for AfdEndpoints ValidateCustomDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_ValidateCustomDomain.json
     */
    /**
     * Sample code: Endpoints_ValidateCustomDomain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsValidateCustomDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdEndpoints()
            .validateCustomDomainWithResponse(
                "RG",
                "profile1",
                "endpoint1",
                new ValidateCustomDomainInput().withHostname("www.someDomain.com"),
                Context.NONE);
    }
}
```
