
import com.azure.core.util.Context;

/** Samples for SubscriptionUsages Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SubscriptionUsageGet.json
     */
    /**
     * Sample code: Get specific subscription usage in the given location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getSpecificSubscriptionUsageInTheGivenLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSubscriptionUsages().getWithResponse("WestUS", "ServerQuota",
            Context.NONE);
    }
}
