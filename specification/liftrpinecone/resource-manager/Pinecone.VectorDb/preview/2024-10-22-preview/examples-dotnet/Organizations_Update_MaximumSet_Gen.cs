using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PineconeVectorDB.Models;
using Azure.ResourceManager.PineconeVectorDB;

// Generated from example definition: 2024-10-22-preview/Organizations_Update_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PineconeVectorDBOrganizationResource created on azure
// for more information of creating PineconeVectorDBOrganizationResource, please refer to the document of PineconeVectorDBOrganizationResource
string subscriptionId = "76a38ef6-c8c1-4f0d-bfe0-00ec782c8077";
string resourceGroupName = "rgopenapi";
string organizationname = "example-organization-name";
ResourceIdentifier pineconeVectorDBOrganizationResourceId = PineconeVectorDBOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationname);
PineconeVectorDBOrganizationResource pineconeVectorDBOrganization = client.GetPineconeVectorDBOrganizationResource(pineconeVectorDBOrganizationResourceId);

// invoke the operation
PineconeVectorDBOrganizationPatch patch = new PineconeVectorDBOrganizationPatch
{
    Tags =
    {
    ["new-tag"] = "new.tag.value"
    },
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("ident573739201")] = new UserAssignedIdentity()
        },
    },
};
PineconeVectorDBOrganizationResource result = await pineconeVectorDBOrganization.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PineconeVectorDBOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
