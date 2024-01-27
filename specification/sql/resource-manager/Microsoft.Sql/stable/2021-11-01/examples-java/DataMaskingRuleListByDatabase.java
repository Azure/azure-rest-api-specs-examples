
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/** Samples for DataMaskingRules ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DataMaskingRuleListByDatabase.json
     */
    /**
     * Sample code: Gets a list of database data masking rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfDatabaseDataMaskingRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDataMaskingRules().listByDatabase("sqlcrudtest-6852",
            "sqlcrudtest-2080", "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, Context.NONE);
    }
}
