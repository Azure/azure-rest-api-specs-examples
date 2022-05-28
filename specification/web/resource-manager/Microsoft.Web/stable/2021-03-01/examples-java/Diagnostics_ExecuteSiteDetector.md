Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Diagnostics ExecuteSiteDetector. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_ExecuteSiteDetector.json
     */
    /**
     * Sample code: Execute site detector.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteDetector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .executeSiteDetectorWithResponse(
                "Sample-WestUSResourceGroup",
                "SampleApp",
                "sitecrashes",
                "availability",
                null,
                null,
                null,
                Context.NONE);
    }
}
```
