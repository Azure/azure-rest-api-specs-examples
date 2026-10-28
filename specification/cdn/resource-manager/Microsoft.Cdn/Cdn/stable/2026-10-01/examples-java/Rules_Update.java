
import com.azure.resourcemanager.cdn.models.DeliveryRuleEdgeActionParameters;
import com.azure.resourcemanager.cdn.models.DeliveryRuleResponseHeaderAction;
import com.azure.resourcemanager.cdn.models.EdgeAction;
import com.azure.resourcemanager.cdn.models.HeaderAction;
import com.azure.resourcemanager.cdn.models.HeaderActionParameters;
import com.azure.resourcemanager.cdn.models.InvocationPoint;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.RuleUpdateParameters;
import java.util.Arrays;

/**
 * Samples for Rules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Rules_Update.json
     */
    /**
     * Sample code: Rules_Update.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void rulesUpdate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRules().update("RG", "profile1", "ruleSet1", "rule1",
            new RuleUpdateParameters().withOrder(1)
                .withActions(Arrays.asList(
                    new DeliveryRuleResponseHeaderAction().withParameters(new HeaderActionParameters()
                        .withHeaderAction(HeaderAction.OVERWRITE).withHeaderName("X-CDN").withValue("MSFT")),
                    new EdgeAction().withParameters(new DeliveryRuleEdgeActionParameters()
                        .withEdgeActionReference(new ResourceReference().withId(
                            "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/edgeActions/edgeAction1"))
                        .withInvocationPoint(InvocationPoint.CLIENT_RESPONSE)))),
            com.azure.core.util.Context.NONE);
    }
}
