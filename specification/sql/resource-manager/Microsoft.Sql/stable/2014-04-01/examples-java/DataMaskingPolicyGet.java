import com.azure.core.util.Context;

/** Samples for DataMaskingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingPolicyGet.json
     */
    /**
     * Sample code: Get data masking policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDataMaskingPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDataMaskingPolicies()
            .getWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", Context.NONE);
    }
}
