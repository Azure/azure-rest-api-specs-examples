Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Annotations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsDelete.json
     */
    /**
     * Sample code: AnnotationsDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void annotationsDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .annotations()
            .deleteWithResponse(
                "my-resource-group", "my-component", "bb820f1b-3110-4a8b-ba2c-8c1129d7eb6a", Context.NONE);
    }
}
```
