using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicySetDefinition.json
// this example is just showing the usage of "PolicySetDefinitions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscription = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SubscriptionPolicySetDefinitionResource
SubscriptionPolicySetDefinitionCollection collection = subscription.GetSubscriptionPolicySetDefinitions();

// invoke the operation
string policySetDefinitionName = "CostManagement";
PolicySetDefinitionData data = new PolicySetDefinitionData()
{
    DisplayName = "Cost Management",
    Description = "Policies to enforce low cost storage SKUs",
    Metadata = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["category"] = "Cost Management"
    }),
    Parameters =
    {
    ["namePrefix"] = new ArmPolicyParameter()
    {
    ParameterType = ArmPolicyParameterType.String,
    DefaultValue = BinaryData.FromString("\"myPrefix\""),
    Metadata = new ParameterDefinitionsValueMetadata()
    {
    DisplayName = "Prefix to enforce on resource names",
    },
    },
    },
    PolicyDefinitions =
    {
    new PolicyDefinitionReference("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1")
    {
    Parameters =
    {
    ["listOfAllowedSKUs"] = new ArmPolicyParameterValue()
    {
    Value = BinaryData.FromObjectAsJson(new object[] { "Standard_GRS", "Standard_LRS" }),
    },
    },
    PolicyDefinitionReferenceId = "Limit_Skus",
    },new PolicyDefinitionReference("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming")
    {
    Parameters =
    {
    ["prefix"] = new ArmPolicyParameterValue()
    {
    Value = BinaryData.FromString("\"[parameters('namePrefix')]\""),
    },
    ["suffix"] = new ArmPolicyParameterValue()
    {
    Value = BinaryData.FromString("\"-LC\""),
    },
    },
    PolicyDefinitionReferenceId = "Resource_Naming",
    }
    },
};
ArmOperation<SubscriptionPolicySetDefinitionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policySetDefinitionName, data);
SubscriptionPolicySetDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicySetDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
