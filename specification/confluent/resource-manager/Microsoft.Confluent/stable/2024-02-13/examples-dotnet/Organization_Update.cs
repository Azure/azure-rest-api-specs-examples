using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Confluent;
using Azure.ResourceManager.Confluent.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_Update.json
// this example is just showing the usage of "Organization_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConfluentOrganizationResource created on azure
// for more information of creating ConfluentOrganizationResource, please refer to the document of ConfluentOrganizationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string organizationName = "myOrganization";
ResourceIdentifier confluentOrganizationResourceId = ConfluentOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
ConfluentOrganizationResource confluentOrganization = client.GetConfluentOrganizationResource(confluentOrganizationResourceId);

// invoke the operation
ConfluentOrganizationPatch patch = new ConfluentOrganizationPatch()
{
    Tags =
    {
    ["client"] = "dev-client",
    ["env"] = "dev",
    },
};
ConfluentOrganizationResource result = await confluentOrganization.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConfluentOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
