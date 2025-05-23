
import com.azure.resourcemanager.applicationinsights.models.WorkbookSharedTypeKind;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workbooks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookAdd.json
     */
    /**
     * Sample code: WorkbookAdd.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookAdd(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.workbooks().define("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2").withRegion("westus")
            .withExistingResourceGroup("my-resource-group")
            .withTags(mapOf("TagSample01", "sample01", "TagSample02", "sample02"))
            .withKind(WorkbookSharedTypeKind.SHARED).withDisplayName("Sample workbook")
            .withSerializedData(
                "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}")
            .withCategory("workbook").withDescription("Sample workbook").withSourceIdParameter(
                "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group")
            .create();
    }

    // Use "Map.of" if available
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
