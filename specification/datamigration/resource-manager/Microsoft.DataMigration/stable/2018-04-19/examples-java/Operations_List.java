/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void operationsList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
