
import com.azure.resourcemanager.sql.fluent.models.DataMaskingRuleInner;
import com.azure.resourcemanager.sql.models.DataMaskingFunction;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/**
 * Samples for DataMaskingRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DataMaskingRuleCreateOrUpdateText.json
     */
    /**
     * Sample code: Create/Update data masking rule for text.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createUpdateDataMaskingRuleForText(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingRules().createOrUpdateWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080",
            "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, "rule1",
            new DataMaskingRuleInner().withSchemaName("dbo").withTableName("Table_1").withColumnName("test1")
                .withMaskingFunction(DataMaskingFunction.TEXT).withPrefixSize("1").withSuffixSize("0")
                .withReplacementString("asdf"),
            com.azure.core.util.Context.NONE);
    }
}
