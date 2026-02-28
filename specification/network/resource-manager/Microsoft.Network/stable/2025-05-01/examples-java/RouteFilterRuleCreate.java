
import com.azure.resourcemanager.network.fluent.models.RouteFilterRuleInner;
import com.azure.resourcemanager.network.models.Access;
import com.azure.resourcemanager.network.models.RouteFilterRuleType;
import java.util.Arrays;

/**
 * Samples for RouteFilterRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteFilterRuleCreate.json
     */
    /**
     * Sample code: RouteFilterRuleCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteFilterRules().createOrUpdate("rg1", "filterName", "ruleName",
            new RouteFilterRuleInner().withAccess(Access.ALLOW).withRouteFilterRuleType(RouteFilterRuleType.COMMUNITY)
                .withCommunities(Arrays.asList("12076:5030", "12076:5040")),
            com.azure.core.util.Context.NONE);
    }
}
