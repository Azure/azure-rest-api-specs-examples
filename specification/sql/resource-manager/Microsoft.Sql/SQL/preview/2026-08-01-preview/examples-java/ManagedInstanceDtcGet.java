
import com.azure.resourcemanager.sql.models.DtcName;

/**
 * Samples for ManagedInstanceDtcs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceDtcGet.json
     */
    /**
     * Sample code: Gets managed instance DTC settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsManagedInstanceDTCSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceDtcs().getWithResponse("testrg", "testinstance", DtcName.CURRENT,
            com.azure.core.util.Context.NONE);
    }
}
