Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.Kind;
import com.azure.resourcemanager.applicationinsights.models.MyWorkbook;

/** Samples for MyWorkbooks Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookUpdate.json
     */
    /**
     * Sample code: WorkbookUpdate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        MyWorkbook resource =
            manager
                .myWorkbooks()
                .getByResourceGroupWithResponse(
                    "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", Context.NONE)
                .getValue();
        resource
            .update()
            .withKind(Kind.USER)
            .withDisplayName("Blah Blah Blah")
            .withSerializedData(
                "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r"
                    + "\\n"
                    + "---\\r"
                    + "\\n"
                    + "\\r"
                    + "\\n"
                    + "Welcome to your new workbook.  This area will display text formatted as markdown.\\r"
                    + "\\n"
                    + "\\r"
                    + "\\n"
                    + "\\r"
                    + "\\n"
                    + "We've included a basic analytics query to get you started. Use the `Edit` button below each"
                    + " section to configure it or add more"
                    + " sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union"
                    + " withsource=TableName *\\n"
                    + "| summarize Count=count() by TableName\\n"
                    + "| render"
                    + " barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}")
            .withVersion("ME")
            .withCategory("workbook")
            .withSourceId(
                "/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens")
            .apply();
    }
}
```
