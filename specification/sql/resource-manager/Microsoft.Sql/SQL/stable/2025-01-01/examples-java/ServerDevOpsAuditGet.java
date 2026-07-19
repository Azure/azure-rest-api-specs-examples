
import com.azure.resourcemanager.sql.models.DevOpsAuditingSettingsName;

/**
 * Samples for ServerDevOpsAuditSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDevOpsAuditGet.json
     */
    /**
     * Sample code: Get a server's DevOps audit settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAServerSDevOpsAuditSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDevOpsAuditSettings().getWithResponse("devAuditTestRG", "devOpsAuditTestSvr",
            DevOpsAuditingSettingsName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
