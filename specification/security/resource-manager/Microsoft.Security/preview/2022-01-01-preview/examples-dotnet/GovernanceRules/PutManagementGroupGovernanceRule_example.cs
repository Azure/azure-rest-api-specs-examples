using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PutManagementGroupGovernanceRule_example.json
// this example is just showing the usage of "GovernanceRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this GovernanceRuleResource
string scope = "providers/Microsoft.Management/managementGroups/contoso";
GovernanceRuleCollection collection = client.GetGovernanceRules(new ResourceIdentifier(scope));

// invoke the operation
string ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
GovernanceRuleData data = new GovernanceRuleData
{
    DisplayName = "Management group rule",
    Description = "A rule for a management group",
    RemediationTimeframe = "7.00:00:00",
    IsGracePeriod = true,
    RulePriority = 200,
    IsDisabled = false,
    RuleType = GovernanceRuleType.Integrated,
    SourceResourceType = GovernanceRuleSourceResourceType.Assessments,
    ExcludedScopes = { "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23" },
    ConditionSets = {BinaryData.FromObjectAsJson(new
    {
    conditions = new object[]
    {
    new Dictionary<string, object>
    {
    ["operator"] = "In",
    ["property"] = "$.AssessmentKey",
    ["value"] = "[\"b1cd27e0-4ecc-4246-939f-49c426d9d72f\", \"fe83f80b-073d-4ccf-93d9-6797eb870201\"]"
    }
    },
    })},
    OwnerSource = new GovernanceRuleOwnerSource
    {
        SourceType = GovernanceRuleOwnerSourceType.Manually,
        Value = "user@contoso.com",
    },
    GovernanceEmailNotification = new GovernanceRuleEmailNotification
    {
        IsManagerEmailNotificationDisabled = true,
        IsOwnerEmailNotificationDisabled = false,
    },
};
ArmOperation<GovernanceRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleId, data);
GovernanceRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GovernanceRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
