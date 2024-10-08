
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/**
 * Samples for DataMaskingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DataMaskingPolicyGet.json
     */
    /**
     * Sample code: Gets the database data masking policies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheDatabaseDataMaskingPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDataMaskingPolicies().getWithResponse("sqlcrudtest-6852",
            "sqlcrudtest-2080", "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
