
import com.azure.resourcemanager.sql.models.GeoBackupPolicyName;

/**
 * Samples for GeoBackupPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GeoBackupPoliciesGet.json
     */
    /**
     * Sample code: Gets the specified Geo backup policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheSpecifiedGeoBackupPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getGeoBackupPolicies().getWithResponse("sqlcrudtest-4799",
            "sqlcrudtest-5961", "testdw", GeoBackupPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
