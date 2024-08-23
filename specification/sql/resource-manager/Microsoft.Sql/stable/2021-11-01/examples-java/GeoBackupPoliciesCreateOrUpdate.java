
import com.azure.resourcemanager.sql.fluent.models.GeoBackupPolicyInner;
import com.azure.resourcemanager.sql.models.GeoBackupPolicyName;
import com.azure.resourcemanager.sql.models.GeoBackupPolicyState;

/**
 * Samples for GeoBackupPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GeoBackupPoliciesCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a database default Geo backup policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateADatabaseDefaultGeoBackupPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getGeoBackupPolicies().createOrUpdateWithResponse(
            "sqlcrudtest-4799", "sqlcrudtest-5961", "testdw", GeoBackupPolicyName.DEFAULT,
            new GeoBackupPolicyInner().withState(GeoBackupPolicyState.ENABLED), com.azure.core.util.Context.NONE);
    }
}
