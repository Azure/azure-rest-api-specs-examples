import com.azure.core.util.Context;

/** Samples for DataMaskingRules ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingRuleList.json
     */
    /**
     * Sample code: List data masking rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDataMaskingRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDataMaskingRules()
            .listByDatabase("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", Context.NONE);
    }
}
