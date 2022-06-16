import com.azure.core.util.Context;
import com.azure.resourcemanager.azurearcdata.models.SqlServerInstance;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlServerInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/UpdateSqlServerInstance.json
     */
    /**
     * Sample code: Updates a SQL Server Instance tags.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void updatesASQLServerInstanceTags(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        SqlServerInstance resource =
            manager
                .sqlServerInstances()
                .getByResourceGroupWithResponse("testrg", "testsqlServerInstance", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("mytag", "myval")).apply();
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
