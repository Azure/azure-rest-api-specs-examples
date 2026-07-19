
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceDtcInner;
import com.azure.resourcemanager.sql.models.DtcName;

/**
 * Samples for ManagedInstanceDtcs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceDtcUpdateEnableDtc.json
     */
    /**
     * Sample code: Updates managed instance DTC settings by enabling DTC.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updatesManagedInstanceDTCSettingsByEnablingDTC(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceDtcs().createOrUpdate("testrg", "testinstance", DtcName.CURRENT,
            new ManagedInstanceDtcInner().withDtcEnabled(true), com.azure.core.util.Context.NONE);
    }
}
