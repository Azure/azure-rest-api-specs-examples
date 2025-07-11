
import com.azure.resourcemanager.network.fluent.models.RouteMapInner;
import com.azure.resourcemanager.network.models.Action;
import com.azure.resourcemanager.network.models.Criterion;
import com.azure.resourcemanager.network.models.NextStep;
import com.azure.resourcemanager.network.models.Parameter;
import com.azure.resourcemanager.network.models.RouteMapActionType;
import com.azure.resourcemanager.network.models.RouteMapMatchCondition;
import com.azure.resourcemanager.network.models.RouteMapRule;
import java.util.Arrays;

/**
 * Samples for RouteMaps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/RouteMapPut.json
     */
    /**
     * Sample code: RouteMapPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeMapPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteMaps().createOrUpdate("rg1", "virtualHub1", "routeMap1",
            new RouteMapInner().withAssociatedInboundConnections(Arrays.asList(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGateway1/expressRouteConnections/exrConn1"))
                .withAssociatedOutboundConnections(Arrays.asList()).withRules(
                    Arrays
                        .asList(new RouteMapRule().withName("rule1")
                            .withMatchCriteria(Arrays.asList(new Criterion()
                                .withRoutePrefix(Arrays.asList("10.0.0.0/8")).withCommunity(Arrays.asList())
                                .withAsPath(Arrays.asList()).withMatchCondition(RouteMapMatchCondition.CONTAINS)))
                            .withActions(Arrays.asList(new Action().withType(RouteMapActionType.ADD)
                                .withParameters(Arrays.asList(new Parameter().withRoutePrefix(Arrays.asList())
                                    .withCommunity(Arrays.asList()).withAsPath(Arrays.asList("22334"))))))
                            .withNextStepIfMatched(NextStep.CONTINUE))),
            com.azure.core.util.Context.NONE);
    }
}
