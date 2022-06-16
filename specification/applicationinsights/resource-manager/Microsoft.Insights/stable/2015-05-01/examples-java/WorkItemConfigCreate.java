import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.WorkItemCreateConfiguration;
import java.util.HashMap;
import java.util.Map;

/** Samples for WorkItemConfigurations Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigCreate.json
     */
    /**
     * Sample code: WorkItemConfigurationsCreate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workItemConfigurationsCreate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workItemConfigurations()
            .createWithResponse(
                "my-resource-group",
                "my-component",
                new WorkItemCreateConfiguration()
                    .withConnectorId("d334e2a4-6733-488e-8645-a9fdc1694f41")
                    .withConnectorDataConfiguration(
                        "{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"d0662b05-439a-4a1b-840b-33a7f8b42ebf\",\"Custom\":\"{\\\"/fields/System.WorkItemType\\\":\\\"Bug\\\",\\\"/fields/System.AreaPath\\\":\\\"todeletefirst\\\",\\\"/fields/System.AssignedTo\\\":\\\"\\\"}\"}")
                    .withValidateOnly(true)
                    .withWorkItemProperties(mapOf()),
                Context.NONE);
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
