
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/**
 * Samples for DataMaskingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DataMaskingPolicyGet.json
     */
    /**
     * Sample code: Gets the database data masking policies.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheDatabaseDataMaskingPolicies(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingPolicies().getWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080",
            "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
