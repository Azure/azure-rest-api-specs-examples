Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WebApps GetFtpAllowed. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicy.json
     */
    /**
     * Sample code: Get FTP Allowed.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFTPAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getFtpAllowedWithResponse("rg", "testSite", Context.NONE);
    }
}
```
