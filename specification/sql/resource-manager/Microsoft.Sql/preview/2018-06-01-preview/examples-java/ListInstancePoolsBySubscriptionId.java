import com.azure.core.util.Context;

/** Samples for InstancePools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ListInstancePoolsBySubscriptionId.json
     */
    /**
     * Sample code: List instance pools in the subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listInstancePoolsInTheSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getInstancePools().list(Context.NONE);
    }
}
