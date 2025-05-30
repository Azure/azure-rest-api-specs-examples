using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyDefinition.json
// this example is just showing the usage of "PolicyDefinitions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionPolicyDefinitionResource created on azure
// for more information of creating SubscriptionPolicyDefinitionResource, please refer to the document of SubscriptionPolicyDefinitionResource
string subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
string policyDefinitionName = "ResourceNaming";
ResourceIdentifier subscriptionPolicyDefinitionResourceId = SubscriptionPolicyDefinitionResource.CreateResourceIdentifier(subscriptionId, policyDefinitionName);
SubscriptionPolicyDefinitionResource subscriptionPolicyDefinition = client.GetSubscriptionPolicyDefinitionResource(subscriptionPolicyDefinitionResourceId);

// invoke the operation
await subscriptionPolicyDefinition.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
