Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.applicationinsights.models.Kind;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workbooks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookAdd.json
     */
    /**
     * Sample code: WorkbookAdd.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookAdd(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .define("deadb33f-8bee-4d3b-a059-9be8dac93960")
            .withRegion("west us")
            .withExistingResourceGroup("my-resource-group")
            .withTags(mapOf("TagSample01", "sample01", "TagSample02", "sample02"))
            .withKind(Kind.SHARED)
            .withDisplayName("tttt")
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
            .withCategory("workbook")
            .withDescription("Sample workbook")
            .withSourceIdParameter("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/MyGroup")
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
