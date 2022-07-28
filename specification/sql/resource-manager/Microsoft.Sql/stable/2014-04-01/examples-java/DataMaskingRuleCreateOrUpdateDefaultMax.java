import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DataMaskingRuleInner;
import com.azure.resourcemanager.sql.models.DataMaskingFunction;
import com.azure.resourcemanager.sql.models.DataMaskingRuleState;

/** Samples for DataMaskingRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingRuleCreateOrUpdateDefaultMax.json
     */
    /**
     * Sample code: Create/Update data masking rule for default max.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateDataMaskingRuleForDefaultMax(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDataMaskingRules()
            .createOrUpdateWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-331",
                "rule1",
                new DataMaskingRuleInner()
                    .withAliasName("nickname")
                    .withRuleState(DataMaskingRuleState.ENABLED)
                    .withSchemaName("dbo")
                    .withTableName("Table_1")
                    .withColumnName("test1")
                    .withMaskingFunction(DataMaskingFunction.DEFAULT),
                Context.NONE);
    }
}
