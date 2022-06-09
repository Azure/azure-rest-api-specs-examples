```java
import com.azure.core.util.Context;

/** Samples for SqlVirtualMachineGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListByResourceGroupSqlVirtualMachineGroup.json
     */
    /**
     * Sample code: Gets all SQL virtual machine groups in a resource group.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsAllSQLVirtualMachineGroupsInAResourceGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachineGroups().listByResourceGroup("testrg", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-sqlvirtualmachine_1.0.0-beta.2/sdk/sqlvirtualmachine/azure-resourcemanager-sqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.
