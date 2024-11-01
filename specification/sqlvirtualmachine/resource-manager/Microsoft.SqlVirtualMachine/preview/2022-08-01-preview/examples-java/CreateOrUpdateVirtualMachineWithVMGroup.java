
import com.azure.resourcemanager.sqlvirtualmachine.models.WsfcDomainCredentials;

/**
 * Samples for SqlVirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * CreateOrUpdateVirtualMachineWithVMGroup.json
     */
    /**
     * Sample code: Creates or updates a SQL virtual machine and joins it to a SQL virtual machine group.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void createsOrUpdatesASQLVirtualMachineAndJoinsItToASQLVirtualMachineGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().define("testvm").withRegion("northeurope").withExistingResourceGroup("testrg")
            .withVirtualMachineResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm2")
            .withWsfcDomainCredentials(
                new WsfcDomainCredentials().withClusterBootstrapAccountPassword("fakeTokenPlaceholder")
                    .withClusterOperatorAccountPassword("fakeTokenPlaceholder")
                    .withSqlServiceAccountPassword("fakeTokenPlaceholder"))
            .withWsfcStaticIp("10.0.0.7").create();
    }
}
