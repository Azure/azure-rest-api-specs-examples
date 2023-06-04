using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_GetResourceScope.json
// this example is just showing the usage of "Remediations_GetAtResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this PolicyRemediationResource
string resourceId = "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceId));
PolicyRemediationCollection collection = client.GetPolicyRemediations(scopeId);

// invoke the operation
string remediationName = "storageRemediation";
bool result = await collection.ExistsAsync(remediationName);

Console.WriteLine($"Succeeded: {result}");
