using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/Attestations_GetResourceScope.json
// this example is just showing the usage of "Attestations_GetAtResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PolicyAttestationResource created on azure
// for more information of creating PolicyAttestationResource, please refer to the document of PolicyAttestationResource
string resourceId = "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM";
string attestationName = "790996e6-9871-4b1f-9cd9-ec42cd6ced1e";
ResourceIdentifier policyAttestationResourceId = PolicyAttestationResource.CreateResourceIdentifier(resourceId, attestationName);
PolicyAttestationResource policyAttestation = client.GetPolicyAttestationResource(policyAttestationResourceId);

// invoke the operation
PolicyAttestationResource result = await policyAttestation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyAttestationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
