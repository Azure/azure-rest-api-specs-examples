using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_CreateResourceScope.json
// this example is just showing the usage of "Attestations_CreateOrUpdateAtResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this PolicyAttestationResource
string resourceId = "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM";
PolicyAttestationCollection collection = client.GetPolicyAttestations(new ResourceIdentifier(resourceId));

// invoke the operation
string attestationName = "790996e6-9871-4b1f-9cd9-ec42cd6ced1e";
PolicyAttestationData data = new PolicyAttestationData(new ResourceIdentifier("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"))
{
    PolicyDefinitionReferenceId = "0b158b46-ff42-4799-8e39-08a5c23b4551",
    ComplianceState = PolicyComplianceState.Compliant,
    ExpireOn = DateTimeOffset.Parse("2021-06-15T00:00:00Z"),
    Owner = "55a32e28-3aa5-4eea-9b5a-4cd85153b966",
    Comments = "This subscription has passed a security audit.",
    Evidence = {new AttestationEvidence
    {
    Description = "The results of the security audit.",
    SourceUri = new Uri("https://gist.github.com/contoso/9573e238762c60166c090ae16b814011"),
    }},
    AssessOn = DateTimeOffset.Parse("2021-06-10T00:00:00Z"),
    Metadata = BinaryData.FromObjectAsJson(new
    {
        departmentId = "NYC-MARKETING-1",
    }),
};
ArmOperation<PolicyAttestationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, attestationName, data);
PolicyAttestationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyAttestationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
