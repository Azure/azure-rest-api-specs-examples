Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.Workbook;

/** Samples for Workbooks Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookUpdate.json
     */
    /**
     * Sample code: WorkbookUpdate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        Workbook resource =
            manager
                .workbooks()
                .getByResourceGroupWithResponse(
                    "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", Context.NONE)
                .getValue();
        resource
            .update()
            .withSourceId(
                "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/Default-Web-WestUS/providers/Microsoft.Web/sites/MyTestApp-CodeLens1")
            .apply();
    }
}
```
