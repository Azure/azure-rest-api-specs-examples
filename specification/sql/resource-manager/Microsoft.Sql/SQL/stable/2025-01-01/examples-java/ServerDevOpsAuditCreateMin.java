
import com.azure.resourcemanager.sql.fluent.models.ServerDevOpsAuditingSettingsInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import com.azure.resourcemanager.sql.models.DevOpsAuditingSettingsName;

/**
 * Samples for ServerDevOpsAuditSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDevOpsAuditCreateMin.json
     */
    /**
     * Sample code: Update a server's DevOps audit settings with minimal input.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateAServerSDevOpsAuditSettingsWithMinimalInput(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDevOpsAuditSettings().createOrUpdate("devAuditTestRG", "devOpsAuditTestSvr",
            DevOpsAuditingSettingsName.DEFAULT,
            new ServerDevOpsAuditingSettingsInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
