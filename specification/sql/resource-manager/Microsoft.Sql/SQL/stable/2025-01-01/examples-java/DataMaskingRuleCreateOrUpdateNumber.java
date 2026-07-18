
import com.azure.resourcemanager.sql.fluent.models.DataMaskingRuleInner;
import com.azure.resourcemanager.sql.models.DataMaskingFunction;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/**
 * Samples for DataMaskingRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DataMaskingRuleCreateOrUpdateNumber.json
     */
    /**
     * Sample code: Create/Update data masking rule for numbers.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createUpdateDataMaskingRuleForNumbers(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingRules().createOrUpdateWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080",
            "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, "rule1",
            new DataMaskingRuleInner().withSchemaName("dbo").withTableName("Table_1").withColumnName("test1")
                .withMaskingFunction(DataMaskingFunction.NUMBER).withNumberFrom("0").withNumberTo("2"),
            com.azure.core.util.Context.NONE);
    }
}
