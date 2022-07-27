import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.GeoBackupPolicyInner;
import com.azure.resourcemanager.sql.models.GeoBackupPolicyName;
import com.azure.resourcemanager.sql.models.GeoBackupPolicyState;

/** Samples for GeoBackupPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/GeoBackupPoliciesCreateOrUpdate.json
     */
    /**
     * Sample code: Update geo backup policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateGeoBackupPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getGeoBackupPolicies()
            .createOrUpdateWithResponse(
                "sqlcrudtest-4799",
                "sqlcrudtest-5961",
                "testdw",
                GeoBackupPolicyName.DEFAULT,
                new GeoBackupPolicyInner().withState(GeoBackupPolicyState.ENABLED),
                Context.NONE);
    }
}
