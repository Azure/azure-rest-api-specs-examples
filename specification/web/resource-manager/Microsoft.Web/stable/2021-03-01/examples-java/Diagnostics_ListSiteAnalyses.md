Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Diagnostics ListSiteAnalysesSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_ListSiteAnalyses.json
     */
    /**
     * Sample code: List App Analyses.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppAnalyses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .listSiteAnalysesSlot(
                "Sample-WestUSResourceGroup", "SampleApp", "availability", "Production", Context.NONE);
    }
}
```
