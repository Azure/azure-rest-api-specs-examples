
import com.azure.core.util.Context;

/** Samples for ServerDevOpsAuditSettings ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDevOpsAuditSettingsList.json
     */
    /**
     * Sample code: List DevOps audit settings of a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDevOpsAuditSettingsOfAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerDevOpsAuditSettings().listByServer("devAuditTestRG",
            "devOpsAuditTestSvr", Context.NONE);
    }
}
