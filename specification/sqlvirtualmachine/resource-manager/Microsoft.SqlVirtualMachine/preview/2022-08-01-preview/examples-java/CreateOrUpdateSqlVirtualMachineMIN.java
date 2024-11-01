
/**
 * Samples for SqlVirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * CreateOrUpdateSqlVirtualMachineMIN.json
     */
    /**
     * Sample code: Creates or updates a SQL virtual machine with min parameters.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void createsOrUpdatesASQLVirtualMachineWithMinParameters(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().define("testvm").withRegion("northeurope").withExistingResourceGroup("testrg")
            .withVirtualMachineResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm")
            .create();
    }
}
