import com.azure.core.util.Context;

/** Samples for ServiceTierAdvisors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServiceTierAdvisorGet.json
     */
    /**
     * Sample code: Get a service tier advisor.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServiceTierAdvisor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServiceTierAdvisors()
            .getWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "Current", Context.NONE);
    }
}
