
import com.azure.resourcemanager.cdn.fluent.models.RuleInner;
import com.azure.resourcemanager.cdn.models.DeliveryRuleEdgeActionParameters;
import com.azure.resourcemanager.cdn.models.DeliveryRuleRequestMethodCondition;
import com.azure.resourcemanager.cdn.models.EdgeAction;
import com.azure.resourcemanager.cdn.models.InvocationPoint;
import com.azure.resourcemanager.cdn.models.RequestMethodMatchConditionParameters;
import com.azure.resourcemanager.cdn.models.RequestMethodMatchConditionParametersMatchValuesItem;
import com.azure.resourcemanager.cdn.models.RequestMethodOperator;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import java.util.Arrays;

/**
 * Samples for Rules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Rules_Create_EdgeAction.json
     */
    /**
     * Sample code: Rules_Create_EdgeAction.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void rulesCreateEdgeAction(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRules().create("RG", "profile1", "ruleSet1", "rule1",
            new RuleInner().withOrder(1)
                .withConditions(Arrays.asList(
                    new DeliveryRuleRequestMethodCondition().withParameters(new RequestMethodMatchConditionParameters()
                        .withOperator(RequestMethodOperator.EQUAL).withNegateCondition(false)
                        .withMatchValues(Arrays.asList(RequestMethodMatchConditionParametersMatchValuesItem.GET)))))
                .withActions(Arrays.asList(new EdgeAction().withParameters(new DeliveryRuleEdgeActionParameters()
                    .withEdgeActionReference(new ResourceReference().withId(
                        "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/edgeActions/edgeAction1"))
                    .withInvocationPoint(InvocationPoint.CLIENT_RESPONSE)))),
            com.azure.core.util.Context.NONE);
    }
}
