
import com.azure.resourcemanager.sql.models.GeoBackupPolicyName;

/**
 * Samples for GeoBackupPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GeoBackupPoliciesGet.json
     */
    /**
     * Sample code: Gets the specified Geo backup policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheSpecifiedGeoBackupPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getGeoBackupPolicies().getWithResponse("sqlcrudtest-4799", "sqlcrudtest-5961", "testdw",
            GeoBackupPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
