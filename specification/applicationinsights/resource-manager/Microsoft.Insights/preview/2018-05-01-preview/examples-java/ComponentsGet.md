Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Components GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsGet.json
     */
    /**
     * Sample code: ComponentGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentGet(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.components().getByResourceGroupWithResponse("my-resource-group", "my-component", Context.NONE);
    }
}
```
