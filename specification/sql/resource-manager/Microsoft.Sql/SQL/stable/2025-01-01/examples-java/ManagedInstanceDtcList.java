
/**
 * Samples for ManagedInstanceDtcs ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceDtcList.json
     */
    /**
     * Sample code: Gets a list of managed instance DTC settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfManagedInstanceDTCSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceDtcs().listByManagedInstance("testrg", "testinstance",
            com.azure.core.util.Context.NONE);
    }
}
