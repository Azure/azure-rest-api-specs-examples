Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-sqlvirtualmachine_1.0.0-beta.2/sdk/sqlvirtualmachine/azure-resourcemanager-sqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/GetSqlVirtualMachine.json
     */
    /**
     * Sample code: Gets a SQL virtual machine.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsASQLVirtualMachine(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().getByResourceGroupWithResponse("testrg", "testvm", null, Context.NONE);
    }
}
```
