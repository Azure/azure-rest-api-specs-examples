using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/WafPolicyCreateOrUpdate.json
// this example is just showing the usage of "WebApplicationFirewallPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebApplicationFirewallPolicyResource created on azure
// for more information of creating WebApplicationFirewallPolicyResource, please refer to the document of WebApplicationFirewallPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string policyName = "Policy1";
ResourceIdentifier webApplicationFirewallPolicyResourceId = WebApplicationFirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, policyName);
WebApplicationFirewallPolicyResource webApplicationFirewallPolicy = client.GetWebApplicationFirewallPolicyResource(webApplicationFirewallPolicyResourceId);

// invoke the operation
WebApplicationFirewallPolicyData data = new WebApplicationFirewallPolicyData
{
    PolicySettings = new PolicySettings
    {
        LogScrubbing = new PolicySettingsLogScrubbing
        {
            State = WebApplicationFirewallScrubbingState.Enabled,
            ScrubbingRules = {new WebApplicationFirewallScrubbingRules(ScrubbingRuleEntryMatchVariable.RequestArgNames, ScrubbingRuleEntryMatchOperator.EqualsValue)
            {
            Selector = "test",
            State = ScrubbingRuleEntryState.Enabled,
            }, new WebApplicationFirewallScrubbingRules(ScrubbingRuleEntryMatchVariable.RequestIPAddress, ScrubbingRuleEntryMatchOperator.EqualsAny)
            {
            State = ScrubbingRuleEntryState.Enabled,
            }},
        },
        JsChallengeCookieExpirationInMins = 100,
    },
    CustomRules = {new WebApplicationFirewallCustomRule(1, WebApplicationFirewallRuleType.MatchRule, new MatchCondition[]
    {
    new MatchCondition(new MatchVariable[]
    {
    new MatchVariable(WebApplicationFirewallMatchVariable.RemoteAddr)
    {
    Selector = null,
    }
    }, WebApplicationFirewallOperator.IPMatch, new string[]{"192.168.1.0/24", "10.0.0.0/24"})
    }, WebApplicationFirewallAction.Block)
    {
    Name = "Rule1",
    }, new WebApplicationFirewallCustomRule(2, WebApplicationFirewallRuleType.MatchRule, new MatchCondition[]
    {
    new MatchCondition(new MatchVariable[]
    {
    new MatchVariable(WebApplicationFirewallMatchVariable.RemoteAddr)
    {
    Selector = null,
    }
    }, WebApplicationFirewallOperator.IPMatch, new string[]{"192.168.1.0/24"}),
    new MatchCondition(new MatchVariable[]
    {
    new MatchVariable(WebApplicationFirewallMatchVariable.RequestHeaders)
    {
    Selector = "UserAgent",
    }
    }, WebApplicationFirewallOperator.Contains, new string[]{"Windows"})
    }, WebApplicationFirewallAction.Block)
    {
    Name = "Rule2",
    }, new WebApplicationFirewallCustomRule(3, WebApplicationFirewallRuleType.RateLimitRule, new MatchCondition[]
    {
    new MatchCondition(new MatchVariable[]
    {
    new MatchVariable(WebApplicationFirewallMatchVariable.RemoteAddr)
    {
    Selector = null,
    }
    }, WebApplicationFirewallOperator.IPMatch, new string[]{"192.168.1.0/24", "10.0.0.0/24"})
    {
    NegationConditon = true,
    }
    }, WebApplicationFirewallAction.Block)
    {
    Name = "RateLimitRule3",
    RateLimitDuration = ApplicationGatewayFirewallRateLimitDuration.OneMin,
    RateLimitThreshold = 10,
    GroupByUserSession = {new GroupByUserSession(new GroupByVariable[]
    {
    new GroupByVariable(ApplicationGatewayFirewallUserSessionVariable.ClientAddr)
    })},
    }, new WebApplicationFirewallCustomRule(4, WebApplicationFirewallRuleType.MatchRule, new MatchCondition[]
    {
    new MatchCondition(new MatchVariable[]
    {
    new MatchVariable(WebApplicationFirewallMatchVariable.RemoteAddr)
    {
    Selector = null,
    }
    }, WebApplicationFirewallOperator.IPMatch, new string[]{"192.168.1.0/24"}),
    new MatchCondition(new MatchVariable[]
    {
    new MatchVariable(WebApplicationFirewallMatchVariable.RequestHeaders)
    {
    Selector = "UserAgent",
    }
    }, WebApplicationFirewallOperator.Contains, new string[]{"Bot"})
    }, WebApplicationFirewallAction.JSChallenge)
    {
    Name = "Rule4",
    }},
    ManagedRules = new ManagedRulesDefinition(new ManagedRuleSet[]
{
new ManagedRuleSet("OWASP", "3.2")
{
RuleGroupOverrides = {new ManagedRuleGroupOverride("REQUEST-931-APPLICATION-ATTACK-RFI")
{
Rules = {new ManagedRuleOverride("931120")
{
State = ManagedRuleEnabledState.Enabled,
Action = RuleMatchActionType.Log,
}, new ManagedRuleOverride("931130")
{
State = ManagedRuleEnabledState.Disabled,
Action = RuleMatchActionType.AnomalyScoring,
}},
}},
},
new ManagedRuleSet("Microsoft_BotManagerRuleSet", "1.0")
{
RuleGroupOverrides = {new ManagedRuleGroupOverride("UnknownBots")
{
Rules = {new ManagedRuleOverride("300700")
{
State = ManagedRuleEnabledState.Enabled,
Action = RuleMatchActionType.JSChallenge,
}},
}},
},
new ManagedRuleSet("Microsoft_HTTPDDoSRuleSet", "1.0")
{
RuleGroupOverrides = {new ManagedRuleGroupOverride("ExcessiveRequests")
{
Rules = {new ManagedRuleOverride("500100")
{
State = ManagedRuleEnabledState.Enabled,
Action = RuleMatchActionType.Block,
Sensitivity = ManagedRuleSensitivityType.High,
}},
}},
}
})
    {
        Exceptions = {new ExceptionEntry(ExceptionEntryMatchVariable.RequestUri, ExceptionEntryValueMatchOperator.Contains)
        {
        Values = {"health", "account/images", "default.aspx"},
        ExceptionManagedRuleSets = {new ExclusionManagedRuleSet("OWASP", "3.2")},
        }, new ExceptionEntry(ExceptionEntryMatchVariable.RequestHeader, ExceptionEntryValueMatchOperator.Contains)
        {
        Values = {"Mozilla/5.0", "Chrome/122.0.0.0"},
        SelectorMatchOperator = ExceptionEntrySelectorMatchOperator.StartsWith,
        Selector = "User-Agent",
        ExceptionManagedRuleSets = {new ExclusionManagedRuleSet("OWASP", "3.2")
        {
        RuleGroups = {new ExclusionManagedRuleGroup("REQUEST-932-APPLICATION-ATTACK-RCE")},
        }},
        }, new ExceptionEntry(ExceptionEntryMatchVariable.RemoteAddr, ExceptionEntryValueMatchOperator.IPMatch)
        {
        Values = {"1.2.3.4", "10.0.0.1/6"},
        ExceptionManagedRuleSets = {new ExclusionManagedRuleSet("Microsoft_BotManagerRuleSet", "1.0")
        {
        RuleGroups = {new ExclusionManagedRuleGroup("BadBots")
        {
        Rules = {new ExclusionManagedRule("100100")},
        }},
        }},
        }},
        Exclusions = {new OwaspCrsExclusionEntry(OwaspCrsExclusionEntryMatchVariable.RequestArgNames, OwaspCrsExclusionEntrySelectorMatchOperator.StartsWith, "hello")
        {
        ExclusionManagedRuleSets = {new ExclusionManagedRuleSet("OWASP", "3.2")
        {
        RuleGroups = {new ExclusionManagedRuleGroup("REQUEST-930-APPLICATION-ATTACK-LFI")
        {
        Rules = {new ExclusionManagedRule("930120")},
        }, new ExclusionManagedRuleGroup("REQUEST-932-APPLICATION-ATTACK-RCE")},
        }},
        }, new OwaspCrsExclusionEntry(OwaspCrsExclusionEntryMatchVariable.RequestArgNames, OwaspCrsExclusionEntrySelectorMatchOperator.EndsWith, "hello")
        {
        ExclusionManagedRuleSets = {new ExclusionManagedRuleSet("OWASP", "3.1")
        {
        RuleGroups = {},
        }},
        }, new OwaspCrsExclusionEntry(OwaspCrsExclusionEntryMatchVariable.RequestArgNames, OwaspCrsExclusionEntrySelectorMatchOperator.StartsWith, "test"), new OwaspCrsExclusionEntry(OwaspCrsExclusionEntryMatchVariable.RequestArgValues, OwaspCrsExclusionEntrySelectorMatchOperator.StartsWith, "test")},
    },
    Location = new AzureLocation("WestUs"),
};
ArmOperation<WebApplicationFirewallPolicyResource> lro = await webApplicationFirewallPolicy.UpdateAsync(WaitUntil.Completed, data);
WebApplicationFirewallPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebApplicationFirewallPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
