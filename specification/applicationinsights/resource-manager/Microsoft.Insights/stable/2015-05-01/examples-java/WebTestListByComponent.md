```java
import com.azure.core.util.Context;

/** Samples for WebTests ListByComponent. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestListByComponent.json
     */
    /**
     * Sample code: webTestListByComponent.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestListByComponent(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().listByComponent("my-component", "my-resource-group", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.
