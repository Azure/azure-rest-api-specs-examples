Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.AfdDomainInner;
import com.azure.resourcemanager.cdn.models.AfdCertificateType;
import com.azure.resourcemanager.cdn.models.AfdDomainHttpsParameters;
import com.azure.resourcemanager.cdn.models.AfdMinimumTlsVersion;
import com.azure.resourcemanager.cdn.models.ResourceReference;

/** Samples for AfdCustomDomains Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_Create.json
     */
    /**
     * Sample code: AFDCustomDomains_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDCustomDomainsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdCustomDomains()
            .create(
                "RG",
                "profile1",
                "domain1",
                new AfdDomainInner()
                    .withHostname("www.someDomain.net")
                    .withTlsSettings(
                        new AfdDomainHttpsParameters()
                            .withCertificateType(AfdCertificateType.MANAGED_CERTIFICATE)
                            .withMinimumTlsVersion(AfdMinimumTlsVersion.TLS12))
                    .withAzureDnsZone(new ResourceReference().withId("")),
                Context.NONE);
    }
}
```
