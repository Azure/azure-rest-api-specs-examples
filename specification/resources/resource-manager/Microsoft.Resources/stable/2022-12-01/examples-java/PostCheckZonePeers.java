
import com.azure.resourcemanager.resources.models.CheckZonePeersRequest;
import java.util.Arrays;

/**
 * Samples for Subscriptions CheckZonePeers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/PostCheckZonePeers.json
     */
    /**
     * Sample code: GetLogicalZoneMapping.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLogicalZoneMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().subscriptionClient().getSubscriptions().checkZonePeersWithResponse(
            "8d65815f-a5b6-402f-9298-045155da7d74", new CheckZonePeersRequest().withLocation("eastus")
                .withSubscriptionIds(Arrays.asList("subscriptions/11111111-1111-1111-1111-111111111111")),
            com.azure.core.util.Context.NONE);
    }
}
