
/**
 * Samples for ServerDevOpsAuditSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDevOpsAuditGet.json
     */
    /**
     * Sample code: Get a server's DevOps audit settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServerSDevOpsAuditSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerDevOpsAuditSettings().getWithResponse("devAuditTestRG",
            "devOpsAuditTestSvr", "default", com.azure.core.util.Context.NONE);
    }
}
