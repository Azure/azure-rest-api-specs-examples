import com.azure.resourcemanager.synapse.models.TransparentDataEncryptionName;

/** Samples for SqlPoolTransparentDataEncryptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolTransparentDataEncryption.json
     */
    /**
     * Sample code: Get transparent data encryption configuration of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getTransparentDataEncryptionConfigurationOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTransparentDataEncryptions()
            .getWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                TransparentDataEncryptionName.CURRENT,
                com.azure.core.util.Context.NONE);
    }
}
