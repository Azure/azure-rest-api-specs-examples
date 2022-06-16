import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.CdnWebApplicationFirewallPolicyInner;
import com.azure.resourcemanager.cdn.models.ActionType;
import com.azure.resourcemanager.cdn.models.CustomRule;
import com.azure.resourcemanager.cdn.models.CustomRuleEnabledState;
import com.azure.resourcemanager.cdn.models.CustomRuleList;
import com.azure.resourcemanager.cdn.models.ManagedRuleEnabledState;
import com.azure.resourcemanager.cdn.models.ManagedRuleGroupOverride;
import com.azure.resourcemanager.cdn.models.ManagedRuleOverride;
import com.azure.resourcemanager.cdn.models.ManagedRuleSet;
import com.azure.resourcemanager.cdn.models.ManagedRuleSetList;
import com.azure.resourcemanager.cdn.models.MatchCondition;
import com.azure.resourcemanager.cdn.models.Operator;
import com.azure.resourcemanager.cdn.models.PolicySettings;
import com.azure.resourcemanager.cdn.models.PolicySettingsDefaultCustomBlockResponseStatusCode;
import com.azure.resourcemanager.cdn.models.RateLimitRule;
import com.azure.resourcemanager.cdn.models.RateLimitRuleList;
import com.azure.resourcemanager.cdn.models.Sku;
import com.azure.resourcemanager.cdn.models.SkuName;
import com.azure.resourcemanager.cdn.models.TransformType;
import com.azure.resourcemanager.cdn.models.WafMatchVariable;
import java.util.Arrays;

/** Samples for Policies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyCreateOrUpdate.json
     */
    /**
     * Sample code: Creates specific policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsSpecificPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getPolicies()
            .createOrUpdate(
                "rg1",
                "MicrosoftCdnWafPolicy",
                new CdnWebApplicationFirewallPolicyInner()
                    .withLocation("WestUs")
                    .withSku(new Sku().withName(SkuName.STANDARD_MICROSOFT))
                    .withPolicySettings(
                        new PolicySettings()
                            .withDefaultRedirectUrl("http://www.bing.com")
                            .withDefaultCustomBlockResponseStatusCode(
                                PolicySettingsDefaultCustomBlockResponseStatusCode.TWO_ZERO_ZERO)
                            .withDefaultCustomBlockResponseBody(
                                "PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="))
                    .withRateLimitRules(
                        new RateLimitRuleList()
                            .withRules(
                                Arrays
                                    .asList(
                                        new RateLimitRule()
                                            .withName("RateLimitRule1")
                                            .withEnabledState(CustomRuleEnabledState.ENABLED)
                                            .withPriority(1)
                                            .withMatchConditions(
                                                Arrays
                                                    .asList(
                                                        new MatchCondition()
                                                            .withMatchVariable(WafMatchVariable.REMOTE_ADDR)
                                                            .withOperator(Operator.IPMATCH)
                                                            .withNegateCondition(false)
                                                            .withMatchValue(
                                                                Arrays.asList("192.168.1.0/24", "10.0.0.0/24"))
                                                            .withTransforms(Arrays.asList())))
                                            .withAction(ActionType.BLOCK)
                                            .withRateLimitThreshold(1000)
                                            .withRateLimitDurationInMinutes(0))))
                    .withCustomRules(
                        new CustomRuleList()
                            .withRules(
                                Arrays
                                    .asList(
                                        new CustomRule()
                                            .withName("CustomRule1")
                                            .withEnabledState(CustomRuleEnabledState.ENABLED)
                                            .withPriority(2)
                                            .withMatchConditions(
                                                Arrays
                                                    .asList(
                                                        new MatchCondition()
                                                            .withMatchVariable(WafMatchVariable.REMOTE_ADDR)
                                                            .withOperator(Operator.GEO_MATCH)
                                                            .withNegateCondition(false)
                                                            .withMatchValue(Arrays.asList("CH"))
                                                            .withTransforms(Arrays.asList()),
                                                        new MatchCondition()
                                                            .withMatchVariable(WafMatchVariable.REQUEST_HEADER)
                                                            .withSelector("UserAgent")
                                                            .withOperator(Operator.CONTAINS)
                                                            .withNegateCondition(false)
                                                            .withMatchValue(Arrays.asList("windows"))
                                                            .withTransforms(Arrays.asList()),
                                                        new MatchCondition()
                                                            .withMatchVariable(WafMatchVariable.QUERY_STRING)
                                                            .withSelector("search")
                                                            .withOperator(Operator.CONTAINS)
                                                            .withNegateCondition(false)
                                                            .withMatchValue(Arrays.asList("<?php", "?>"))
                                                            .withTransforms(
                                                                Arrays
                                                                    .asList(
                                                                        TransformType.URL_DECODE,
                                                                        TransformType.LOWERCASE))))
                                            .withAction(ActionType.BLOCK))))
                    .withManagedRules(
                        new ManagedRuleSetList()
                            .withManagedRuleSets(
                                Arrays
                                    .asList(
                                        new ManagedRuleSet()
                                            .withRuleSetType("DefaultRuleSet")
                                            .withRuleSetVersion("preview-1.0")
                                            .withRuleGroupOverrides(
                                                Arrays
                                                    .asList(
                                                        new ManagedRuleGroupOverride()
                                                            .withRuleGroupName("Group1")
                                                            .withRules(
                                                                Arrays
                                                                    .asList(
                                                                        new ManagedRuleOverride()
                                                                            .withRuleId("GROUP1-0001")
                                                                            .withEnabledState(
                                                                                ManagedRuleEnabledState.ENABLED)
                                                                            .withAction(ActionType.REDIRECT),
                                                                        new ManagedRuleOverride()
                                                                            .withRuleId("GROUP1-0002")
                                                                            .withEnabledState(
                                                                                ManagedRuleEnabledState
                                                                                    .DISABLED)))))))),
                Context.NONE);
    }
}
