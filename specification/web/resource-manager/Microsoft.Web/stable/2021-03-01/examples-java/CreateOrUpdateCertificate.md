Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.CertificateInner;
import java.util.Arrays;

/** Samples for Certificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateCertificate.json
     */
    /**
     * Sample code: Create Or Update Certificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getCertificates()
            .createOrUpdateWithResponse(
                "testrg123",
                "testc6282",
                new CertificateInner()
                    .withLocation("East US")
                    .withPassword("<password>")
                    .withHostNames(Arrays.asList("ServerCert")),
                Context.NONE);
    }
}
```
