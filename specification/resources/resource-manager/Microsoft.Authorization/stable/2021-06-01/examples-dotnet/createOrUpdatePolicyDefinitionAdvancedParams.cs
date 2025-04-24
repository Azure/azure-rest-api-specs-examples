using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAdvancedParams.json
// this example is just showing the usage of "PolicyDefinitions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionPolicyDefinitionResource created on azure
// for more information of creating SubscriptionPolicyDefinitionResource, please refer to the document of SubscriptionPolicyDefinitionResource
string subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
string policyDefinitionName = "EventHubDiagnosticLogs";
ResourceIdentifier subscriptionPolicyDefinitionResourceId = SubscriptionPolicyDefinitionResource.CreateResourceIdentifier(subscriptionId, policyDefinitionName);
SubscriptionPolicyDefinitionResource subscriptionPolicyDefinition = client.GetSubscriptionPolicyDefinitionResource(subscriptionPolicyDefinitionResourceId);

// invoke the operation
PolicyDefinitionData data = new PolicyDefinitionData
{
    Mode = "Indexed",
    DisplayName = "Event Hubs should have diagnostic logging enabled",
    Description = "Audit enabling of logs and retain them up to a year. This enables recreation of activity trails for investigation purposes when a security incident occurs or your network is compromised",
    PolicyRule = BinaryData.FromObjectAsJson(new Dictionary<string, object>
    {
        ["if"] = new
        {
            equals = "Microsoft.EventHub/namespaces",
            field = "type",
        },
        ["then"] = new
        {
            effect = "AuditIfNotExists",
            details = new
            {
                type = "Microsoft.Insights/diagnosticSettings",
                existenceCondition = new
                {
                    allOf = new object[]
{
new
{
equals = "true",
field = "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled",
},
new
{
equals = "[parameters('requiredRetentionDays')]",
field = "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days",
}
},
                },
            },
        }
    }),
    Metadata = BinaryData.FromObjectAsJson(new
    {
        category = "Event Hub",
    }),
    Parameters =
    {
    ["requiredRetentionDays"] = new ArmPolicyParameter
    {
    ParameterType = ArmPolicyParameterType.Integer,
    AllowedValues = {BinaryData.FromObjectAsJson("0"), BinaryData.FromObjectAsJson("30"), BinaryData.FromObjectAsJson("90"), BinaryData.FromObjectAsJson("180"), BinaryData.FromObjectAsJson("365")},
    DefaultValue = BinaryData.FromObjectAsJson("365"),
    Metadata = new ParameterDefinitionsValueMetadata
    {
    DisplayName = "Required retention (days)",
    Description = "The required diagnostic logs retention in days",
    },
    }
    },
};
ArmOperation<SubscriptionPolicyDefinitionResource> lro = await subscriptionPolicyDefinition.UpdateAsync(WaitUntil.Completed, data);
SubscriptionPolicyDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
