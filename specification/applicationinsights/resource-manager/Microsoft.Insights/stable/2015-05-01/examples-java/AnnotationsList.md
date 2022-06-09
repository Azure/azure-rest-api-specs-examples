```java
import com.azure.core.util.Context;

/** Samples for Annotations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsList.json
     */
    /**
     * Sample code: AnnotationsList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void annotationsList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .annotations()
            .list(
                "my-resource-group",
                "my-component",
                "2018-02-05T00%3A30%3A00.000Z",
                "2018-02-06T00%3A33A00.000Z",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.
