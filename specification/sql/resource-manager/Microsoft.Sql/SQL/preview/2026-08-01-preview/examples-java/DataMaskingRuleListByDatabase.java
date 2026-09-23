
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/**
 * Samples for DataMaskingRules ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DataMaskingRuleListByDatabase.json
     */
    /**
     * Sample code: Gets a list of database data masking rules.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfDatabaseDataMaskingRules(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingRules().listByDatabase("sqlcrudtest-6852", "sqlcrudtest-2080",
            "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, null, com.azure.core.util.Context.NONE);
    }
}
