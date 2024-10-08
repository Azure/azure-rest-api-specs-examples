
import com.azure.resourcemanager.sql.fluent.models.ServerDevOpsAuditingSettingsInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import java.util.UUID;

/**
 * Samples for ServerDevOpsAuditSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDevOpsAuditCreateMax.json
     */
    /**
     * Sample code: Update a server's DevOps audit settings with all params.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateAServerSDevOpsAuditSettingsWithAllParams(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerDevOpsAuditSettings().createOrUpdate("devAuditTestRG",
            "devOpsAuditTestSvr", "default",
            new ServerDevOpsAuditingSettingsInner().withIsAzureMonitorTargetEnabled(true)
                .withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder").withStorageAccountSubscriptionId(
                    UUID.fromString("00000000-1234-0000-5678-000000000000")),
            com.azure.core.util.Context.NONE);
    }
}
