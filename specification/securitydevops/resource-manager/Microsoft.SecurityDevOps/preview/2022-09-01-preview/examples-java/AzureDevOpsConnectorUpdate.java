import com.azure.core.util.Context;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsConnector;
import java.util.HashMap;
import java.util.Map;

/** Samples for AzureDevOpsConnector Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorUpdate.json
     */
    /**
     * Sample code: AzureDevOpsConnector_Update.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsConnectorUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        AzureDevOpsConnector resource =
            manager
                .azureDevOpsConnectors()
                .getByResourceGroupWithResponse("westusrg", "testconnector", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("client", "dev-client", "env", "dev")).apply();
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
