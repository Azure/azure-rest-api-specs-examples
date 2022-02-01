Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Components List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsList.json
     */
    /**
     * Sample code: ComponentsList.json.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentsListJson(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.components().list(Context.NONE);
    }
}
```
