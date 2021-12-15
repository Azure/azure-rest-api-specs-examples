Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Annotations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsGet.json
     */
    /**
     * Sample code: AnnotationsGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void annotationsGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .annotations()
            .getWithResponse("my-resource-group", "my-component", "444e2c08-274a-4bbb-a89e-d77bb720f44a", Context.NONE);
    }
}
```
