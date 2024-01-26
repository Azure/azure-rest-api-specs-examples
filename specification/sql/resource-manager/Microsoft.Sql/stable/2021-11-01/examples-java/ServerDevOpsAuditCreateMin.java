
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerDevOpsAuditingSettingsInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for ServerDevOpsAuditSettings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDevOpsAuditCreateMin.json
     */
    /**
     * Sample code: Update a server's DevOps audit settings with minimal input.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateAServerSDevOpsAuditSettingsWithMinimalInput(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerDevOpsAuditSettings().createOrUpdate("devAuditTestRG",
            "devOpsAuditTestSvr", "default",
            new ServerDevOpsAuditingSettingsInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            Context.NONE);
    }
}
