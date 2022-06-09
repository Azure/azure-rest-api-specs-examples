```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_Get.json
     */
    /**
     * Sample code: Get a private endpoint.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAPrivateEndpoint(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.privateEndpoints().getWithResponse("sjrg", "testcluster", "testpe", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
