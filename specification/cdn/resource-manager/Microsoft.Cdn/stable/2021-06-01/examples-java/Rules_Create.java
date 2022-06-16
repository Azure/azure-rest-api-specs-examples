import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.RuleInner;
import com.azure.resourcemanager.cdn.models.DeliveryRuleRequestMethodCondition;
import com.azure.resourcemanager.cdn.models.DeliveryRuleResponseHeaderAction;
import com.azure.resourcemanager.cdn.models.HeaderAction;
import com.azure.resourcemanager.cdn.models.HeaderActionParameters;
import com.azure.resourcemanager.cdn.models.RequestMethodMatchConditionParameters;
import com.azure.resourcemanager.cdn.models.RequestMethodMatchConditionParametersMatchValuesItem;
import com.azure.resourcemanager.cdn.models.RequestMethodOperator;
import java.util.Arrays;

/** Samples for Rules Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Create.json
     */
    /**
     * Sample code: Rules_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRules()
            .create(
                "RG",
                "profile1",
                "ruleSet1",
                "rule1",
                new RuleInner()
                    .withOrder(1)
                    .withConditions(
                        Arrays
                            .asList(
                                new DeliveryRuleRequestMethodCondition()
                                    .withParameters(
                                        new RequestMethodMatchConditionParameters()
                                            .withOperator(RequestMethodOperator.EQUAL)
                                            .withNegateCondition(false)
                                            .withMatchValues(
                                                Arrays
                                                    .asList(
                                                        RequestMethodMatchConditionParametersMatchValuesItem.GET)))))
                    .withActions(
                        Arrays
                            .asList(
                                new DeliveryRuleResponseHeaderAction()
                                    .withParameters(
                                        new HeaderActionParameters()
                                            .withHeaderAction(HeaderAction.OVERWRITE)
                                            .withHeaderName("X-CDN")
                                            .withValue("MSFT")))),
                Context.NONE);
    }
}
