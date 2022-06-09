```java
import com.azure.core.util.Context;

/** Samples for Transformations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Transformation_Get.json
     */
    /**
     * Sample code: Get a transformation.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getATransformation(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.transformations().getWithResponse("sjrg4423", "sj8374", "transformation952", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
