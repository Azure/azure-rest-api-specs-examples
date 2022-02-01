Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.CategoryType;

/** Samples for Workbooks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbooksList.json
     */
    /**
     * Sample code: WorkbooksList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbooksList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .listByResourceGroup(
                "my-resource-group",
                CategoryType.WORKBOOK,
                null,
                "/subscriptions/ad2f1a83-caac-4e21-9d2a-9ca3f87105e2/resourceGroups/Default-Web-WestUS/providers/Microsoft.Web/sites/MyTestApp-CodeLens1",
                null,
                Context.NONE);
    }
}
```
