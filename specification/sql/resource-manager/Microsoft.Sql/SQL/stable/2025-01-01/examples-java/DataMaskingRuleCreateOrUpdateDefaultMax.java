
import com.azure.resourcemanager.sql.fluent.models.DataMaskingRuleInner;
import com.azure.resourcemanager.sql.models.DataMaskingFunction;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;
import com.azure.resourcemanager.sql.models.DataMaskingRuleState;

/**
 * Samples for DataMaskingRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DataMaskingRuleCreateOrUpdateDefaultMax.json
     */
    /**
     * Sample code: Create/Update data masking rule for default max.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createUpdateDataMaskingRuleForDefaultMax(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingRules().createOrUpdateWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080",
            "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, "rule1",
            new DataMaskingRuleInner().withRuleState(DataMaskingRuleState.ENABLED).withSchemaName("dbo")
                .withTableName("Table_1").withColumnName("test1").withAliasName("nickname")
                .withMaskingFunction(DataMaskingFunction.DEFAULT),
            com.azure.core.util.Context.NONE);
    }
}
