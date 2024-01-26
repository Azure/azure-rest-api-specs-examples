
import com.azure.core.util.Context;

/** Samples for SubscriptionUsages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SubscriptionUsageListByLocation.json
     */
    /**
     * Sample code: List subscription usages in the given location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSubscriptionUsagesInTheGivenLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSubscriptionUsages().listByLocation("WestUS", Context.NONE);
    }
}
