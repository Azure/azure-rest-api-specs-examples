using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicySetDefinition.json
// this example is just showing the usage of "PolicySetDefinitions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionPolicySetDefinitionResource created on azure
// for more information of creating SubscriptionPolicySetDefinitionResource, please refer to the document of SubscriptionPolicySetDefinitionResource
string subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
string policySetDefinitionName = "CostManagement";
ResourceIdentifier subscriptionPolicySetDefinitionResourceId = SubscriptionPolicySetDefinitionResource.CreateResourceIdentifier(subscriptionId, policySetDefinitionName);
SubscriptionPolicySetDefinitionResource subscriptionPolicySetDefinition = client.GetSubscriptionPolicySetDefinitionResource(subscriptionPolicySetDefinitionResourceId);

// invoke the operation
SubscriptionPolicySetDefinitionResource result = await subscriptionPolicySetDefinition.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicySetDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
