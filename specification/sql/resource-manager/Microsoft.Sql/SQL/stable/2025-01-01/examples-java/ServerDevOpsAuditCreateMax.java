
import com.azure.resourcemanager.sql.fluent.models.ServerDevOpsAuditingSettingsInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import com.azure.resourcemanager.sql.models.DevOpsAuditingSettingsName;
import java.util.UUID;

/**
 * Samples for ServerDevOpsAuditSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDevOpsAuditCreateMax.json
     */
    /**
     * Sample code: Update a server's DevOps audit settings with all params.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateAServerSDevOpsAuditSettingsWithAllParams(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDevOpsAuditSettings().createOrUpdate("devAuditTestRG", "devOpsAuditTestSvr",
            DevOpsAuditingSettingsName.DEFAULT,
            new ServerDevOpsAuditingSettingsInner().withIsAzureMonitorTargetEnabled(true)
                .withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder").withStorageAccountSubscriptionId(
                    UUID.fromString("00000000-1234-0000-5678-000000000000")),
            com.azure.core.util.Context.NONE);
    }
}
