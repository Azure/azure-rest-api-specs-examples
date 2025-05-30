
import com.azure.resourcemanager.network.fluent.models.WebApplicationFirewallPolicyInner;
import com.azure.resourcemanager.network.models.ActionType;
import com.azure.resourcemanager.network.models.ApplicationGatewayFirewallRateLimitDuration;
import com.azure.resourcemanager.network.models.ApplicationGatewayFirewallUserSessionVariable;
import com.azure.resourcemanager.network.models.ExceptionEntry;
import com.azure.resourcemanager.network.models.ExceptionEntryMatchVariable;
import com.azure.resourcemanager.network.models.ExceptionEntrySelectorMatchOperator;
import com.azure.resourcemanager.network.models.ExceptionEntryValueMatchOperator;
import com.azure.resourcemanager.network.models.ExclusionManagedRule;
import com.azure.resourcemanager.network.models.ExclusionManagedRuleGroup;
import com.azure.resourcemanager.network.models.ExclusionManagedRuleSet;
import com.azure.resourcemanager.network.models.GroupByUserSession;
import com.azure.resourcemanager.network.models.GroupByVariable;
import com.azure.resourcemanager.network.models.ManagedRuleEnabledState;
import com.azure.resourcemanager.network.models.ManagedRuleGroupOverride;
import com.azure.resourcemanager.network.models.ManagedRuleOverride;
import com.azure.resourcemanager.network.models.ManagedRuleSet;
import com.azure.resourcemanager.network.models.ManagedRulesDefinition;
import com.azure.resourcemanager.network.models.MatchCondition;
import com.azure.resourcemanager.network.models.MatchVariable;
import com.azure.resourcemanager.network.models.OwaspCrsExclusionEntry;
import com.azure.resourcemanager.network.models.OwaspCrsExclusionEntryMatchVariable;
import com.azure.resourcemanager.network.models.OwaspCrsExclusionEntrySelectorMatchOperator;
import com.azure.resourcemanager.network.models.PolicySettings;
import com.azure.resourcemanager.network.models.PolicySettingsLogScrubbing;
import com.azure.resourcemanager.network.models.ScrubbingRuleEntryMatchOperator;
import com.azure.resourcemanager.network.models.ScrubbingRuleEntryMatchVariable;
import com.azure.resourcemanager.network.models.ScrubbingRuleEntryState;
import com.azure.resourcemanager.network.models.SensitivityType;
import com.azure.resourcemanager.network.models.WebApplicationFirewallAction;
import com.azure.resourcemanager.network.models.WebApplicationFirewallCustomRule;
import com.azure.resourcemanager.network.models.WebApplicationFirewallMatchVariable;
import com.azure.resourcemanager.network.models.WebApplicationFirewallOperator;
import com.azure.resourcemanager.network.models.WebApplicationFirewallRuleType;
import com.azure.resourcemanager.network.models.WebApplicationFirewallScrubbingRules;
import com.azure.resourcemanager.network.models.WebApplicationFirewallScrubbingState;
import java.util.Arrays;

/**
 * Samples for WebApplicationFirewallPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/WafPolicyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a WAF policy within a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsOrUpdatesAWAFPolicyWithinAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getWebApplicationFirewallPolicies()
            .createOrUpdateWithResponse("rg1", "Policy1",
                new WebApplicationFirewallPolicyInner()
                    .withLocation(
                        "WestUs")
                    .withPolicySettings(
                        new PolicySettings()
                            .withLogScrubbing(
                                new PolicySettingsLogScrubbing().withState(WebApplicationFirewallScrubbingState.ENABLED)
                                    .withScrubbingRules(
                                        Arrays
                                            .asList(
                                                new WebApplicationFirewallScrubbingRules()
                                                    .withMatchVariable(
                                                        ScrubbingRuleEntryMatchVariable.REQUEST_ARG_NAMES)
                                                    .withSelectorMatchOperator(ScrubbingRuleEntryMatchOperator.EQUALS)
                                                    .withSelector("test").withState(ScrubbingRuleEntryState.ENABLED),
                                                new WebApplicationFirewallScrubbingRules()
                                                    .withMatchVariable(
                                                        ScrubbingRuleEntryMatchVariable.REQUEST_IPADDRESS)
                                                    .withSelectorMatchOperator(
                                                        ScrubbingRuleEntryMatchOperator.EQUALS_ANY)
                                                    .withState(ScrubbingRuleEntryState.ENABLED))))
                            .withJsChallengeCookieExpirationInMins(100))
                    .withCustomRules(Arrays.asList(
                        new WebApplicationFirewallCustomRule().withName("Rule1").withPriority(1)
                            .withRuleType(WebApplicationFirewallRuleType.MATCH_RULE)
                            .withMatchConditions(Arrays.asList(new MatchCondition()
                                .withMatchVariables(Arrays.asList(new MatchVariable()
                                    .withVariableName(WebApplicationFirewallMatchVariable.REMOTE_ADDR)))
                                .withOperator(WebApplicationFirewallOperator.IPMATCH)
                                .withMatchValues(Arrays.asList("192.168.1.0/24", "10.0.0.0/24"))))
                            .withAction(WebApplicationFirewallAction.BLOCK),
                        new WebApplicationFirewallCustomRule()
                            .withName("Rule2").withPriority(2).withRuleType(
                                WebApplicationFirewallRuleType.MATCH_RULE)
                            .withMatchConditions(Arrays.asList(
                                new MatchCondition()
                                    .withMatchVariables(Arrays.asList(new MatchVariable()
                                        .withVariableName(WebApplicationFirewallMatchVariable.REMOTE_ADDR)))
                                    .withOperator(WebApplicationFirewallOperator.IPMATCH)
                                    .withMatchValues(Arrays.asList("192.168.1.0/24")),
                                new MatchCondition()
                                    .withMatchVariables(Arrays.asList(new MatchVariable()
                                        .withVariableName(WebApplicationFirewallMatchVariable.REQUEST_HEADERS)
                                        .withSelector("UserAgent")))
                                    .withOperator(
                                        WebApplicationFirewallOperator.CONTAINS)
                                    .withMatchValues(Arrays.asList("Windows"))))
                            .withAction(WebApplicationFirewallAction.BLOCK),
                        new WebApplicationFirewallCustomRule().withName("RateLimitRule3").withPriority(3)
                            .withRateLimitDuration(ApplicationGatewayFirewallRateLimitDuration.ONE_MIN)
                            .withRateLimitThreshold(10).withRuleType(WebApplicationFirewallRuleType.RATE_LIMIT_RULE)
                            .withMatchConditions(Arrays.asList(new MatchCondition()
                                .withMatchVariables(Arrays.asList(new MatchVariable()
                                    .withVariableName(WebApplicationFirewallMatchVariable.REMOTE_ADDR)))
                                .withOperator(WebApplicationFirewallOperator.IPMATCH).withNegationConditon(true)
                                .withMatchValues(Arrays.asList("192.168.1.0/24", "10.0.0.0/24"))))
                            .withGroupByUserSession(Arrays.asList(
                                new GroupByUserSession().withGroupByVariables(Arrays.asList(new GroupByVariable()
                                    .withVariableName(ApplicationGatewayFirewallUserSessionVariable.CLIENT_ADDR)))))
                            .withAction(WebApplicationFirewallAction.BLOCK),
                        new WebApplicationFirewallCustomRule()
                            .withName("Rule4").withPriority(4).withRuleType(
                                WebApplicationFirewallRuleType.MATCH_RULE)
                            .withMatchConditions(Arrays.asList(
                                new MatchCondition()
                                    .withMatchVariables(Arrays.asList(new MatchVariable()
                                        .withVariableName(WebApplicationFirewallMatchVariable.REMOTE_ADDR)))
                                    .withOperator(WebApplicationFirewallOperator.IPMATCH)
                                    .withMatchValues(Arrays.asList("192.168.1.0/24")),
                                new MatchCondition()
                                    .withMatchVariables(Arrays.asList(new MatchVariable()
                                        .withVariableName(WebApplicationFirewallMatchVariable.REQUEST_HEADERS)
                                        .withSelector("UserAgent")))
                                    .withOperator(WebApplicationFirewallOperator.CONTAINS)
                                    .withMatchValues(Arrays.asList("Bot"))))
                            .withAction(WebApplicationFirewallAction.JSCHALLENGE)))
                    .withManagedRules(
                        new ManagedRulesDefinition()
                            .withExceptions(Arrays.asList(
                                new ExceptionEntry().withMatchVariable(ExceptionEntryMatchVariable.REQUEST_URI)
                                    .withValues(Arrays.asList("health", "account/images", "default.aspx"))
                                    .withValueMatchOperator(ExceptionEntryValueMatchOperator.CONTAINS)
                                    .withExceptionManagedRuleSets(Arrays.asList(new ExclusionManagedRuleSet()
                                        .withRuleSetType("OWASP").withRuleSetVersion("3.2"))),
                                new ExceptionEntry().withMatchVariable(ExceptionEntryMatchVariable.REQUEST_HEADER)
                                    .withValues(Arrays.asList("Mozilla/5.0", "Chrome/122.0.0.0"))
                                    .withValueMatchOperator(ExceptionEntryValueMatchOperator.CONTAINS)
                                    .withSelectorMatchOperator(ExceptionEntrySelectorMatchOperator.STARTS_WITH)
                                    .withSelector("User-Agent").withExceptionManagedRuleSets(
                                        Arrays.asList(new ExclusionManagedRuleSet().withRuleSetType("OWASP")
                                            .withRuleSetVersion("3.2")
                                            .withRuleGroups(Arrays
                                                .asList(new ExclusionManagedRuleGroup()
                                                    .withRuleGroupName("REQUEST-932-APPLICATION-ATTACK-RCE"))))),
                                new ExceptionEntry().withMatchVariable(ExceptionEntryMatchVariable.REMOTE_ADDR)
                                    .withValues(Arrays.asList("1.2.3.4", "10.0.0.1/6"))
                                    .withValueMatchOperator(ExceptionEntryValueMatchOperator.IPMATCH)
                                    .withExceptionManagedRuleSets(
                                        Arrays.asList(new ExclusionManagedRuleSet()
                                            .withRuleSetType("Microsoft_BotManagerRuleSet").withRuleSetVersion("1.0")
                                            .withRuleGroups(Arrays
                                                .asList(new ExclusionManagedRuleGroup().withRuleGroupName("BadBots")
                                                    .withRules(Arrays
                                                        .asList(new ExclusionManagedRule().withRuleId("100100")))))))))
                            .withExclusions(
                                Arrays
                                    .asList(
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_NAMES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.STARTS_WITH)
                                            .withSelector("hello").withExclusionManagedRuleSets(
                                                Arrays.asList(new ExclusionManagedRuleSet().withRuleSetType("OWASP")
                                                    .withRuleSetVersion("3.2")
                                                    .withRuleGroups(Arrays.asList(
                                                        new ExclusionManagedRuleGroup()
                                                            .withRuleGroupName("REQUEST-930-APPLICATION-ATTACK-LFI")
                                                            .withRules(Arrays.asList(
                                                                new ExclusionManagedRule().withRuleId("930120"))),
                                                        new ExclusionManagedRuleGroup().withRuleGroupName(
                                                            "REQUEST-932-APPLICATION-ATTACK-RCE"))))),
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_NAMES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.ENDS_WITH)
                                            .withSelector("hello").withExclusionManagedRuleSets(
                                                Arrays.asList(new ExclusionManagedRuleSet().withRuleSetType("OWASP")
                                                    .withRuleSetVersion("3.1").withRuleGroups(Arrays.asList()))),
                                        new OwaspCrsExclusionEntry()
                                            .withMatchVariable(OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_NAMES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.STARTS_WITH)
                                            .withSelector("test"),
                                        new OwaspCrsExclusionEntry().withMatchVariable(
                                            OwaspCrsExclusionEntryMatchVariable.REQUEST_ARG_VALUES)
                                            .withSelectorMatchOperator(
                                                OwaspCrsExclusionEntrySelectorMatchOperator.STARTS_WITH)
                                            .withSelector("test")))
                            .withManagedRuleSets(
                                Arrays
                                    .asList(
                                        new ManagedRuleSet().withRuleSetType("OWASP").withRuleSetVersion("3.2")
                                            .withRuleGroupOverrides(
                                                Arrays.asList(new ManagedRuleGroupOverride()
                                                    .withRuleGroupName("REQUEST-931-APPLICATION-ATTACK-RFI")
                                                    .withRules(Arrays.asList(
                                                        new ManagedRuleOverride().withRuleId("931120").withState(
                                                            ManagedRuleEnabledState.ENABLED).withAction(ActionType.LOG),
                                                        new ManagedRuleOverride().withRuleId("931130")
                                                            .withState(ManagedRuleEnabledState.DISABLED)
                                                            .withAction(ActionType.ANOMALY_SCORING))))),
                                        new ManagedRuleSet().withRuleSetType("Microsoft_BotManagerRuleSet")
                                            .withRuleSetVersion("1.0")
                                            .withRuleGroupOverrides(Arrays
                                                .asList(new ManagedRuleGroupOverride().withRuleGroupName("UnknownBots")
                                                    .withRules(Arrays.asList(new ManagedRuleOverride()
                                                        .withRuleId("300700").withState(ManagedRuleEnabledState.ENABLED)
                                                        .withAction(ActionType.JSCHALLENGE))))),
                                        new ManagedRuleSet().withRuleSetType("Microsoft_HTTPDDoSRuleSet")
                                            .withRuleSetVersion("1.0")
                                            .withRuleGroupOverrides(Arrays.asList(
                                                new ManagedRuleGroupOverride().withRuleGroupName("ExcessiveRequests")
                                                    .withRules(Arrays.asList(new ManagedRuleOverride()
                                                        .withRuleId("500100").withState(ManagedRuleEnabledState.ENABLED)
                                                        .withAction(ActionType.BLOCK)
                                                        .withSensitivity(SensitivityType.HIGH)))))))),
                com.azure.core.util.Context.NONE);
    }
}
