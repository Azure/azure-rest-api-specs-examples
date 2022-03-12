Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Diagnostics ListSiteAnalyses. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_ListSiteAnalysesSlot.json
     */
    /**
     * Sample code: List App Slot Analyses.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSlotAnalyses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .listSiteAnalyses("Sample-WestUSResourceGroup", "SampleApp", "availability", Context.NONE);
    }
}
```
