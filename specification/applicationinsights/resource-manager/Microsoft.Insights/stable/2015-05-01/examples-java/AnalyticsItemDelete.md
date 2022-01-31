Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.ItemScopePath;

/** Samples for AnalyticsItems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemDelete.json
     */
    /**
     * Sample code: AnalyticsItemDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void analyticsItemDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .analyticsItems()
            .deleteWithResponse(
                "my-resource-group",
                "my-component",
                ItemScopePath.ANALYTICS_ITEMS,
                "3466c160-4a10-4df8-afdf-0007f3f6dee5",
                null,
                Context.NONE);
    }
}
```
