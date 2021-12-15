Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.fluent.models.ApplicationInsightsComponentAnalyticsItemInner;
import com.azure.resourcemanager.applicationinsights.models.ItemScope;
import com.azure.resourcemanager.applicationinsights.models.ItemScopePath;
import com.azure.resourcemanager.applicationinsights.models.ItemType;

/** Samples for AnalyticsItems Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemPut.json
     */
    /**
     * Sample code: AnalyticsItemPut.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void analyticsItemPut(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .analyticsItems()
            .putWithResponse(
                "my-resource-group",
                "my-component",
                ItemScopePath.ANALYTICS_ITEMS,
                new ApplicationInsightsComponentAnalyticsItemInner()
                    .withName("Exceptions - New in the last 24 hours")
                    .withContent(
                        "let newExceptionsTimeRange = 1d;\n"
                            + "let timeRangeToCheckBefore = 7d;\n"
                            + "exceptions\n"
                            + "| where timestamp < ago(timeRangeToCheckBefore)\n"
                            + "| summarize count() by problemId\n"
                            + "| join kind= rightanti (\n"
                            + "exceptions\n"
                            + "| where timestamp >= ago(newExceptionsTimeRange)\n"
                            + "| extend stack = tostring(details[0].rawStack)\n"
                            + "| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp),"
                            + " any(stack) by problemId  \n"
                            + ") on problemId \n"
                            + "| order by  count_ desc\n")
                    .withScope(ItemScope.SHARED)
                    .withType(ItemType.QUERY),
                null,
                Context.NONE);
    }
}
```
