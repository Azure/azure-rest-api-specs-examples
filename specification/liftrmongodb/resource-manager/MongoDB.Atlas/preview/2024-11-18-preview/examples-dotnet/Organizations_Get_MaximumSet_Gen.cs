using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MongoDBAtlas.Models;
using Azure.ResourceManager.MongoDBAtlas;

// Generated from example definition: 2024-11-18-preview/Organizations_Get_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBAtlasOrganizationResource created on azure
// for more information of creating MongoDBAtlasOrganizationResource, please refer to the document of MongoDBAtlasOrganizationResource
string subscriptionId = "4AFC1287-D389-4265-B2D4-59B96A45CACC";
string resourceGroupName = "rgopenapi";
string organizationName = "U.1-:7";
ResourceIdentifier mongoDBAtlasOrganizationResourceId = MongoDBAtlasOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
MongoDBAtlasOrganizationResource mongoDBAtlasOrganization = client.GetMongoDBAtlasOrganizationResource(mongoDBAtlasOrganizationResourceId);

// invoke the operation
MongoDBAtlasOrganizationResource result = await mongoDBAtlasOrganization.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoDBAtlasOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
