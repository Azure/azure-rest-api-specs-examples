import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines Redeploy. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/RedeploySqlVirtualMachine.json
     */
    /**
     * Sample code: Uninstalls and reinstalls the SQL Iaas Extension.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void uninstallsAndReinstallsTheSQLIaasExtension(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().redeploy("testrg", "testvm", Context.NONE);
    }
}
