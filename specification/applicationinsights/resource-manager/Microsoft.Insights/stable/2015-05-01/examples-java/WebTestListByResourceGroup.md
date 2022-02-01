Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WebTests ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestListByResourceGroup.json
     */
    /**
     * Sample code: webTestListByResourceGroup.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestListByResourceGroup(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().listByResourceGroup("my-resource-group", Context.NONE);
    }
}
```
