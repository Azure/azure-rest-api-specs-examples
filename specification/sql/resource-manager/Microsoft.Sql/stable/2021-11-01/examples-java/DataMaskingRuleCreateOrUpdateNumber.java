
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DataMaskingRuleInner;
import com.azure.resourcemanager.sql.models.DataMaskingFunction;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;

/** Samples for DataMaskingRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DataMaskingRuleCreateOrUpdateNumber.
     * json
     */
    /**
     * Sample code: Create/Update data masking rule for numbers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateDataMaskingRuleForNumbers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDataMaskingRules()
            .createOrUpdateWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331",
                DataMaskingPolicyName.DEFAULT, "rule1",
                new DataMaskingRuleInner().withSchemaName("dbo").withTableName("Table_1").withColumnName("test1")
                    .withMaskingFunction(DataMaskingFunction.NUMBER).withNumberFrom("0").withNumberTo("2"),
                Context.NONE);
    }
}
