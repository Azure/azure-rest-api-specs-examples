
import com.azure.core.util.Context;

/** Samples for ExtendedServerBlobAuditingPolicies ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerExtendedAuditingSettingsList.
     * json
     */
    /**
     * Sample code: List extended auditing settings of a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExtendedAuditingSettingsOfAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getExtendedServerBlobAuditingPolicies()
            .listByServer("blobauditingtest-4799", "blobauditingtest-6440", Context.NONE);
    }
}
