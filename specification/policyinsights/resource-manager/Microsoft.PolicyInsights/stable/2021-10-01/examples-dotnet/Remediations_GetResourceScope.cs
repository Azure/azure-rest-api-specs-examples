using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_GetResourceScope.json
// this example is just showing the usage of "Remediations_GetAtResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PolicyRemediationResource created on azure
// for more information of creating PolicyRemediationResource, please refer to the document of PolicyRemediationResource
string resourceId = "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1";
string remediationName = "storageRemediation";
ResourceIdentifier policyRemediationResourceId = PolicyRemediationResource.CreateResourceIdentifier(resourceId, remediationName);
PolicyRemediationResource policyRemediation = client.GetPolicyRemediationResource(policyRemediationResourceId);

// invoke the operation
PolicyRemediationResource result = await policyRemediation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyRemediationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
