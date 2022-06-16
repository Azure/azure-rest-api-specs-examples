import com.azure.core.util.Context;

/** Samples for SqlPoolTransparentDataEncryptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolTransparentDataEncryptionList.json
     */
    /**
     * Sample code: Get list of transparent data encryption configurations of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getListOfTransparentDataEncryptionConfigurationsOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTransparentDataEncryptions()
            .list("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", Context.NONE);
    }
}
