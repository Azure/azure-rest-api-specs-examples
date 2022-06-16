import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.WebApplicationFirewallPolicyInner;
import com.azure.resourcemanager.network.models.ExclusionManagedRule;
import com.azure.resourcemanager.network.models.ExclusionManagedRuleGroup;
import com.azure.resourcemanager.network.models.ExclusionManagedRuleSet;
import com.azure.resourcemanager.network.models.ManagedRuleSet;
import com.azure.resourcemanager.network.models.ManagedRulesDefinition;
import com.azure.resourcemanager.network.models.MatchCondition;
import com.azure.resourcemanager.network.models.MatchVariable;
import com.azure.resourcemanager.network.models.OwaspCrsExclusionEntry;
import com.azure.resourcemanager.network.models.OwaspCrsExclusionEntryMatchVariable;
import com.azure.resourcemanager.network.models.OwaspCrsExclusionEntrySelectorMatchOperator;
import com.azure.resourcemanager.network.models.WebApplicationFirewallAction;
import com.azure.resourcemanager.network.models.WebApplicationFirewallCustomRule;
import com.azure.resourcemanager.network.models.WebApplicationFirewallMatchVariable;
import com.azure.resourcemanager.network.models.WebApplicationFirewallOperator;
import com.azure.resourcemanager.network.models.WebApplicationFirewallRuleType;
import java.util.Arrays;

/** Samples for WebApplicationFirewallPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a WAF policy within a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesAWAFPolicyWithinAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getWebApplicationFirewallPolicies()
            .createOrUpdateWithResponse(
                "rg1",
                "Policy1",
                new WebApplicationFirewallPolicyInner()
                    .withLocation("WestUs")
                    .withCustomRules(
                        Arrays
                            .asList(
                                new WebApplicationFirewallCustomRule()
                                    .withName("Rule1")
                                    .withPriority(1)
                                    .withRuleType(WebApplicationFirewallRuleType.MATCH_RULE)
                                    .withMatchConditions(
                                        Arrays
                                            .asList(
                                                new MatchCondition()
                                                    .withMatchVariables(
                                                        Arrays
                                                            .asList(
                                                                new MatchVariable()
                                                                    .withVariableName(
                                                                        WebApplicationFirewallMatchVariable
                                                                            .REMOTE_ADDR)))
                                                    .withOperator(WebApplicationFirewallOperator.IPMATCH)
                                                    .withMatchValues(Arrays.asList("192.168.1.0/24", "10.0.0.0/24"))))
                                    .withAction(WebApplicationFirewallAction.BLOCK),
                                new WebApplicationFirewallCustomRule()
                                    .withName("Rule2")
                                    .withPriority(2)
                                    .withRuleType(WebApplicationFirewallRuleType.MATCH_RULE)
                                    .withMatchConditions(
                                        Arrays
                                            .asList(
                                                new MatchCondition()
                                                    .withMatchVariables(
                                                        Arrays
                                                            .asList(
                                                                new MatchVariable()
                                                                    .withVariableName(
                                                                        WebApplicationFirewallMatchVariable
                                                                            .REMOTE_ADDR)))
                                                    .withOperator(WebApplicationFirewallOperator.IPMATCH)
                                                    .withMatchValues(Arrays.asList("192.168.1.0/24")),
                                                new MatchCondition()
                                                    .withMatchVariables(
                                                        Arrays
                                                            .asList(
                                                                new MatchVariable()
                                                                    .withVariableName(
                                                                        WebApplicationFirewallMatchVariable
                                                                            .REQUEST_HEADERS)
                                                                    .withSelector("UserAgent")))
                                                    .withOperator(WebApplicationFirewallOperator.CONTAINS)
                                                    .withMatchValues(Arrays.asList("Windows"))))
                                    .withAction(WebApplicationFirewallAction.BLOCK)))
                    .withManagedRules(
                        new ManagedRulesDefinition()
                            .withExclusions(
                                Arrays
                                    .asList(
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_NAMES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.STARTS_WITH)
                                            .withSelector("hello")
                                            .withExclusionManagedRuleSets(
                                                Arrays
                                                    .asList(
                                                        new ExclusionManagedRuleSet()
                                                            .withRuleSetType("OWASP")
                                                            .withRuleSetVersion("3.2")
                                                            .withRuleGroups(
                                                                Arrays
                                                                    .asList(
                                                                        new ExclusionManagedRuleGroup()
                                                                            .withRuleGroupName(
                                                                                "REQUEST-930-APPLICATION-ATTACK-LFI")
                                                                            .withRules(
                                                                                Arrays
                                                                                    .asList(
                                                                                        new ExclusionManagedRule()
                                                                                            .withRuleId("930120"))),
                                                                        new ExclusionManagedRuleGroup()
                                                                            .withRuleGroupName(
                                                                                "REQUEST-932-APPLICATION-ATTACK-RCE"))))),
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_NAMES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.ENDS_WITH)
                                            .withSelector("hello")
                                            .withExclusionManagedRuleSets(
                                                Arrays
                                                    .asList(
                                                        new ExclusionManagedRuleSet()
                                                            .withRuleSetType("OWASP")
                                                            .withRuleSetVersion("3.1")
                                                            .withRuleGroups(Arrays.asList()))),
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_NAMES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.STARTS_WITH)
                                            .withSelector("test"),
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_VALUES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.STARTS_WITH)
                                            .withSelector("test")))
                            .withManagedRuleSets(
                                Arrays
                                    .asList(new ManagedRuleSet().withRuleSetType("OWASP").withRuleSetVersion("3.2")))),
                Context.NONE);
    }
}
