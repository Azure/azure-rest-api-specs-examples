import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListOperation.json
     */
    /**
     * Sample code: Lists all of the available SQL Virtual Machine Rest API operations.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void listsAllOfTheAvailableSQLVirtualMachineRestAPIOperations(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.operations().list(Context.NONE);
    }
}
