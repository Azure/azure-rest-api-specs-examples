import com.azure.core.util.Context;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlVirtualMachine;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlVirtualMachines Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/UpdateSqlVirtualMachine.json
     */
    /**
     * Sample code: Updates a SQL virtual machine tags.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void updatesASQLVirtualMachineTags(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        SqlVirtualMachine resource =
            manager
                .sqlVirtualMachines()
                .getByResourceGroupWithResponse("testrg", "testvm", null, Context.NONE)
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
