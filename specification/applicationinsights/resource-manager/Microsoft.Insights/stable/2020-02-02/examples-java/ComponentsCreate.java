import com.azure.resourcemanager.applicationinsights.models.ApplicationType;
import com.azure.resourcemanager.applicationinsights.models.FlowType;
import com.azure.resourcemanager.applicationinsights.models.RequestSource;
import java.util.HashMap;
import java.util.Map;

/** Samples for Components CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsCreate.json
     */
    /**
     * Sample code: ComponentCreate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentCreate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .components()
            .define("my-component")
            .withRegion("South Central US")
            .withExistingResourceGroup("my-resource-group")
            .withKind("web")
            .withApplicationType(ApplicationType.WEB)
            .withFlowType(FlowType.BLUEFIELD)
            .withRequestSource(RequestSource.REST)
            .withWorkspaceResourceId(
                "/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace")
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
