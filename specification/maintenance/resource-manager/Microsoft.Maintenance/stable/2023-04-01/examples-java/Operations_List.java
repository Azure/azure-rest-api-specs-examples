/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void operationsList(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
