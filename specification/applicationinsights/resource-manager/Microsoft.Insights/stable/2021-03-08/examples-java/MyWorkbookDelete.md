Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for MyWorkbooks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookDelete.json
     */
    /**
     * Sample code: WorkbookDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .myWorkbooks()
            .deleteWithResponse("my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", Context.NONE);
    }
}
```
