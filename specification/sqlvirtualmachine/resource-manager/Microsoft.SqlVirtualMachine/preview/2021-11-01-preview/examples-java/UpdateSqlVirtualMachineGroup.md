Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-sqlvirtualmachine_1.0.0-beta.2/sdk/sqlvirtualmachine/azure-resourcemanager-sqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlVirtualMachineGroup;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlVirtualMachineGroups Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/UpdateSqlVirtualMachineGroup.json
     */
    /**
     * Sample code: Updates a SQL virtual machine group tags.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void updatesASQLVirtualMachineGroupTags(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        SqlVirtualMachineGroup resource =
            manager
                .sqlVirtualMachineGroups()
                .getByResourceGroupWithResponse("testrg", "testvmgroup", Context.NONE)
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
```
