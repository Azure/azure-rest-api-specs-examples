import com.azure.resourcemanager.loganalytics.models.WorkspaceSku;
import com.azure.resourcemanager.loganalytics.models.WorkspaceSkuNameEnum;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesCreate.json
     */
    /**
     * Sample code: WorkspacesCreate.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .workspaces()
            .define("oiautorest6685")
            .withRegion("australiasoutheast")
            .withExistingResourceGroup("oiautorest6685")
            .withTags(mapOf("tag1", "val1"))
            .withSku(new WorkspaceSku().withName(WorkspaceSkuNameEnum.PER_GB2018))
            .withRetentionInDays(30)
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
